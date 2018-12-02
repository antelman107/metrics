package service

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/antelman107/metrics/echo"
)

type httpServer struct {
	e *echo.Echo

	stopWg   sync.WaitGroup
	stopMux  sync.Mutex
	stopChan chan struct{}
}

func NewHttpServer(e *echo.Echo, logger *zap.Logger) Service {
	var w = &httpServer{
		e:        e,
		stopChan: make(chan struct{}),
	}

	close(w.stopChan)

	return w
}

func (s *httpServer) Start(_ int) (err error) {
	s.stopMux.Lock()
	defer s.stopMux.Unlock()

	select {
	case <-s.stopChan:
	default:
		return ErrAlreadyStarted
	}

	// Start server
	s.stopWg.Add(1)
	go func() {
		defer s.stopWg.Done()

		if err := s.e.Start(":8081"); err != nil {
			s.e.Logger.Info("shutting down the server")
		}
	}()

	s.stopChan = make(chan struct{})

	return nil
}

func (s *httpServer) Stop() error {
	s.stopMux.Lock()
	defer s.stopMux.Unlock()

	select {
	case <-s.stopChan:
		return ErrNotStarted
	default:
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		s.e.Logger.Error(err)
	}

	close(s.stopChan)
	s.stopWg.Wait()

	return nil
}

func (s *httpServer) Done() <-chan struct{} {
	return s.stopChan
}
