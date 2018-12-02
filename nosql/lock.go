package nosql

import (
	"time"
)

type LockInterface interface {
	One(key string, keyDuration time.Duration) (bool, error)
	Iterative(key string, keyDuration time.Duration, waitDuration time.Duration) error
	Release(key string) (bool, error)
}
