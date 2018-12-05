package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/cmd/app/queue"
	"github.com/antelman107/metrics/nosql"
)

type requester struct {
	metricRepo domain.MetricRepository
	subscriber nosql.ListSubscriberInterface
	logger     *zap.Logger

	stopWg   sync.WaitGroup
	stopMux  sync.Mutex
	stopChan chan struct{}
}

func NewRequester(
	metricRepo domain.MetricRepository,
	subscriber nosql.ListSubscriberInterface,
	logger *zap.Logger,
) Service {
	var w = &requester{
		metricRepo: metricRepo,
		subscriber: subscriber,
		logger:     logger.Named("requester"),
		stopChan:   make(chan struct{}),
	}

	close(w.stopChan)

	return w
}

func (s *requester) Start(concurrency int) (err error) {
	s.stopMux.Lock()
	defer s.stopMux.Unlock()

	select {
	case <-s.stopChan:
	default:
		return ErrAlreadyStarted
	}

	s.stopChan = make(chan struct{})

	for i := 0; i < concurrency; i++ {
		s.stopWg.Add(1)

		go s.watch()
	}

	return nil
}

func (s *requester) Stop() error {
	s.stopMux.Lock()
	defer s.stopMux.Unlock()

	select {
	case <-s.stopChan:
		return ErrNotStarted
	default:
	}

	close(s.stopChan)
	s.stopWg.Wait()

	return nil
}

func (s *requester) Done() <-chan struct{} {
	return s.stopChan
}

func (s *requester) watch() {
	defer s.stopWg.Done()

	var (
		data   []byte
		site   *queue.Site
		metric *domain.Metric
		err    error
		delay  = time.Millisecond * 100
		logger = s.logger

		tr = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}
		client    = &http.Client{Transport: tr, Timeout: time.Second * 10}
		resp      *http.Response
		timeStart time.Time
	)

	logger.Info("Start goroutine")

	for {
		select {
		case <-s.stopChan:
			logger.Info("Stop goroutine")

			return

		case <-time.After(delay):
		}

		data, err = s.subscriber.Get()
		if err != nil {
			logger.Error("Can't get request task", zap.Error(err))

			return
		}

		if data == nil {
			continue
		}

		err = json.Unmarshal(data, &site)
		if err != nil {
			logger.Error("Can't unmarshal task", zap.Error(err))

			continue
		}

		logger.Info("Begin request site", zap.String("site", site.Url))

		timeStart = time.Now()
		// @TODO normalize url
		resp, err = client.Get("http://" + site.Url)
		if err != nil {
			logger.Error("Error request site", zap.Error(err), zap.String("site", site.Url))
		}

		metric = &domain.Metric{
			SiteID:    site.ID,
			IsSuccess: true,
			Duration:  time.Now().Sub(timeStart).Nanoseconds(),
		}
		if err != nil {
			metric.Error = err.Error()
			metric.IsSuccess = false
		} else {
			if resp.StatusCode >= http.StatusBadRequest {
				metric.Error = fmt.Sprintf("error code %d", resp.StatusCode)
				metric.IsSuccess = false
			}
		}

		err = resp.Body.Close()
		if err != nil {
			logger.Error("Error response body closing", zap.Error(err))
		}

		err = s.metricRepo.Add(metric)
		if err != nil {
			logger.Error("Error add metric", zap.Error(err), zap.String("site", site.Url), zap.Reflect("metric", metric))

			continue
		}
	}
}
