package nosql

import (
	"fmt"
	"sync"
)

type PubSubSubscriberMock struct {
	sync.Mutex

	mode         int
	channel      string
	isSubscribed bool
	data         [][]byte
}

func (s *PubSubSubscriberMock) Receive() ([]byte, error) {
	if !s.isSubscribed {
		return nil, fmt.Errorf("Before receive quotes you must subscribe of them.")
	}
	s.Lock()
	defer s.Unlock()

	if len(s.data) == 0 {
		return nil, fmt.Errorf("No data")
	}

	value := s.data[0]

	s.data = append(s.data[:0], s.data[1:]...)
	return value, nil
}

func (s *PubSubSubscriberMock) Subscribe(key string) error {
	s.isSubscribed = true

	return nil
}

func (s *PubSubSubscriberMock) Unsubscribe() error {
	s.isSubscribed = false
	return nil
}

func (s *PubSubSubscriberMock) SetData(data [][]byte) {
	s.data = data
}

func NewPubSubSubscriberMock(data [][]byte) *PubSubSubscriberMock {
	return &PubSubSubscriberMock{
		data: data,
	}
}
