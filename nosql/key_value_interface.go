package nosql

import "time"

type (
	KeyValueInterface interface {
		Get(key string) ([]byte, error)
		Set(key string, ttl time.Duration, data []byte) error
		Del(key string) error
		IsExist(key string) (bool, error)
	}
)
