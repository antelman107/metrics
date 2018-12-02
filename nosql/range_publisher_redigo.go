package nosql

import (
	"github.com/gomodule/redigo/redis"
)

type RangePublisherRedigo struct {
	key       string
	redisPool *redis.Pool
}

func (sh *RangePublisherRedigo) Publish(score int64, data []byte) error {
	var conn redis.Conn
	if conn = sh.redisPool.Get(); conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()

	if _, err := conn.Do("ZADD", sh.key, score, data); err != nil {
		return err
	}

	return nil
}

func NewRangePublisherRedigo(
	key string,
	redisPool *redis.Pool,
) *RangePublisherRedigo {
	return &RangePublisherRedigo{
		key:       key,
		redisPool: redisPool,
	}
}
