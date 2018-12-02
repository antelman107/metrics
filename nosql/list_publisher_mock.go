package nosql

import (
	"sync"
)

type ListPublisherMock struct {
	sync.Mutex

	key     string
	storage map[string][][]byte
}

func (sh *ListPublisherMock) Publish(data []byte) error {
	sh.Lock()
	defer sh.Unlock()

	if sh.storage[sh.key] == nil {
		sh.storage[sh.key] = make([][]byte, 0)
	}
	sh.storage[sh.key] = append(sh.storage[sh.key], data)
	return nil
}

func (sh *ListPublisherMock) GetStorage() map[string][][]byte {
	sh.Lock()
	defer sh.Unlock()

	return sh.storage
}

func NewListPublisherMock(
	key string,
) *ListPublisherMock {
	return &ListPublisherMock{
		key:     key,
		storage: make(map[string][][]byte),
	}
}
