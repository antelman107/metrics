package nosql

import (
	"sync"
)

type CounterMockStorage map[string]int64

type CounterMock struct {
	sync.Mutex

	storage CounterMockStorage
	key     string
}

func (cim *CounterMock) GetValue() (int64, error) {
	cim.Lock()
	defer cim.Unlock()

	_, ok := cim.storage[cim.key]
	if !ok {
		cim.storage[cim.key] = 1
		return int64(1), nil
	}
	cim.storage[cim.key] += 1

	return cim.storage[cim.key], nil
}

func NewCounterMock(key string) *CounterMock {
	return &CounterMock{
		key:     key,
		storage: make(CounterMockStorage),
	}
}
