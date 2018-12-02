package nosql

import (
	"sync"
	"time"
)

type lockMockStorage map[string]int

type LockMock struct {
	sync.Mutex
	storage lockMockStorage
}

func (l *LockMock) One(key string, keyDuration time.Duration) (bool, error) {
	l.Lock()
	defer l.Unlock()

	if _, ok := l.storage[key]; ok {
		return false, nil
	}

	l.storage[key] = 1

	if keyDuration > 0 {
		go func(key string) {
			<-time.NewTimer(keyDuration).C
			l.Lock()
			defer l.Unlock()

			if _, ok := l.storage[key]; ok {
				delete(l.storage, key)
			}
		}(key)
	}

	return true, nil
}

func (l *LockMock) Iterative(key string, keyDuration time.Duration, waitDuration time.Duration) error {
	for {
		res, err := l.One(key, keyDuration)
		if err != nil {
			return err
		} else {
			if !res {
				if waitDuration > 0 {
					time.Sleep(waitDuration)
				}
				continue
			}

			return nil
		}
	}
}

func (l *LockMock) Release(key string) error {
	l.Lock()
	defer l.Unlock()

	if _, ok := l.storage[key]; ok {
		delete(l.storage, key)
	}
	return nil
}

func NewLockMock() *LockMock {
	return &LockMock{
		storage: make(lockMockStorage, 0),
	}
}
