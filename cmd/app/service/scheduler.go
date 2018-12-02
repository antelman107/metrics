package service

import (
	"encoding/json"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/cmd/app/queue"
	"github.com/antelman107/metrics/nosql"
)

type scheduler struct {
	siteRepo  domain.SiteRepository
	publisher nosql.ListPublisherInterface
	logger    *zap.Logger

	stopWg   sync.WaitGroup
	stopMux  sync.Mutex
	stopChan chan struct{}
}

func NewScheduler(
	siteRepo domain.SiteRepository,
	publisher nosql.ListPublisherInterface,
	logger *zap.Logger,
) Service {
	var w = &scheduler{
		siteRepo:  siteRepo,
		publisher: publisher,
		logger:    logger.Named("scheduler"),
		stopChan:  make(chan struct{}),
	}

	close(w.stopChan)

	return w
}

func (s *scheduler) Start(concurrency int) (err error) {
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

func (s *scheduler) Stop() error {
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

func (s *scheduler) Done() <-chan struct{} {
	return s.stopChan
}

func (s *scheduler) watch() {
	defer s.stopWg.Done()

	var (
		site   *domain.Site
		sites  domain.Sites
		err    error
		delay  = time.Minute
		logger = s.logger
	)

	logger.Info("Start goroutine")

	for {
		select {
		case <-s.stopChan:
			logger.Info("Stop goroutine")

			return

		case <-time.After(delay):
		}

		sites, err = s.siteRepo.GetAll()
		if err != nil {
			logger.Error("Can't get sites", zap.Error(err))

			continue
		}

		for _, site = range sites {
			js, err := json.Marshal(queue.Site{
				ID:  site.ID,
				Url: site.Url,
			})
			if err != nil {
				logger.Error("Can't marshall site", zap.Error(err))

				continue
			}

			err = s.publisher.Publish(js)
			if err != nil {
				logger.Error("Can't publish site", zap.Error(err))

				continue
			}

			logger.Info("Successfully scheduled", zap.String("site", site.Url))
		}

	}
}
