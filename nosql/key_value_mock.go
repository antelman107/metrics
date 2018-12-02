package nosql

import (
	"time"

	"sync"
)

type (
	KeyValueMock struct {
		sync.Mutex

		storage map[string][]byte
	}
)

func NewKeyValueMock() *KeyValueMock {
	return &KeyValueMock{
		storage: make(map[string][]byte),
	}
}

func (r *KeyValueMock) Get(key string) ([]byte, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.storage[key]; !ok {
		return nil, nil
	}
	val := r.storage[key]

	return val, nil
}

func (r *KeyValueMock) Set(key string, ttl time.Duration, data []byte) error {
	r.Lock()
	defer r.Unlock()

	r.storage[key] = data

	if ttl >= 0 {
		go func() {
			time.Sleep(ttl)

			if _, ok := r.storage[key]; !ok {
				return
			}

			r.Lock()
			defer r.Unlock()
			delete(r.storage, key)
		}()
	}

	return nil
}

func (r *KeyValueMock) Del(key string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.storage[key]; !ok {
		return nil
	}

	delete(r.storage, key)

	return nil
}

func (r *KeyValueMock) IsExist(key string) (bool, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.storage[key]; !ok {
		return false, nil
	}

	return true, nil
}
