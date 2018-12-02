package nosql

import (
	"sync"
)

type PubSubPublisherMock struct {
	sync.Mutex

	isSaveNeeded bool
	storage      map[string][][]byte
}

func (p *PubSubPublisherMock) Publish(key string, data []byte) error {
	if p.isSaveNeeded {
		p.Lock()
		defer p.Unlock()

		p.storage[key] = append(p.storage[key], data)
	}
	return nil
}

func (p *PubSubPublisherMock) GetStorage() map[string][][]byte {
	p.Lock()
	defer p.Unlock()

	return p.storage
}

func NewPubSubPublisherMock(isSaveNeeded bool) *PubSubPublisherMock {
	return &PubSubPublisherMock{
		isSaveNeeded: isSaveNeeded,
		storage:      make(map[string][][]byte),
	}
}
