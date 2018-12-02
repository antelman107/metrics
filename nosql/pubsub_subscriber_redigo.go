package nosql

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type PubSubSubscriberRedigo struct {
	channel    string
	redisPool  *redis.Pool
	pubSubConn *redis.PubSubConn

	wg       *sync.WaitGroup
	stopChan chan struct{}
	stopOnce *sync.Once

	timeout time.Duration
}

func (s *PubSubSubscriberRedigo) Receive() ([]byte, error) {
	s.wg.Add(1)
	defer s.wg.Done()

	if s.pubSubConn == nil {
		return nil, fmt.Errorf("Before receive quotes you must subscribe of them.")
	}

	for {
		select {
		case <-s.stopChan:
			return nil, ErrSubscriberFinished

		default:
			switch n := s.pubSubConn.ReceiveWithTimeout(s.timeout).(type) {
			case redis.Message:
				return n.Data, nil
			case error:
				return nil, n
			}
		}
	}

	return nil, nil
}

func (s *PubSubSubscriberRedigo) Subscribe(key string) (err error) {
	conn := s.redisPool.Get()
	if conn.Err() != nil {
		return conn.Err()
	}

	s.pubSubConn = &redis.PubSubConn{Conn: conn}

	s.wg = &sync.WaitGroup{}

	if err = s.pubSubConn.Subscribe(key); err != nil {
		return err
	}

	s.stopChan = make(chan struct{})
	s.stopOnce = &sync.Once{}

	return nil
}

func (s *PubSubSubscriberRedigo) Unsubscribe() (err error) {
	if err := s.pubSubConn.Unsubscribe(s.channel); err != nil {
		return err
	}

	// Close stop chan
	close(s.stopChan)

	// Wait for subscriber to finish
	s.wg.Wait()

	err = s.pubSubConn.Close()
	if err != nil {
		return err
	}

	s.pubSubConn = nil

	s.wg = nil
	s.stopChan = nil
	s.stopOnce = nil

	return nil
}

func NewPubSubSubscriberRedigo(redisClient *redis.Pool, timeout time.Duration) *PubSubSubscriberRedigo {
	return &PubSubSubscriberRedigo{
		redisPool: redisClient,
		timeout:   timeout,
	}
}
