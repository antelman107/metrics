package nosql

import (
	"github.com/gomodule/redigo/redis"
)

type CounterRedigo struct {
	key       string
	redisPool *redis.Pool
}

func (rc *CounterRedigo) GetValue() (int64, error) {
	var conn redis.Conn
	if conn = rc.redisPool.Get(); conn.Err() != nil {
		return 0, conn.Err()
	}
	defer conn.Close()

	return redis.Int64(conn.Do("INCR", rc.key))
}

func NewCounterRedigo(key string, redisPool *redis.Pool) *CounterRedigo {
	return &CounterRedigo{
		key:       key,
		redisPool: redisPool,
	}
}
