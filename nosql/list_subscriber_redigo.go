package nosql

import (
	"github.com/gomodule/redigo/redis"
)

type ListSubscriberRedigo struct {
	key       string
	redisPool *redis.Pool
}

func (sh *ListSubscriberRedigo) Get() ([]byte, error) {
	var conn redis.Conn
	if conn = sh.redisPool.Get(); conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("LPOP", sh.key))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}

		return nil, err
	}
	return data, err
}

func NewListSubscriberRedigo(
	key string,
	redisPool *redis.Pool,
) *ListSubscriberRedigo {
	return &ListSubscriberRedigo{
		key:       key,
		redisPool: redisPool,
	}
}
