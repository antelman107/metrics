package nosql

import (
	"github.com/gomodule/redigo/redis"
)

type ListPublisherRedigo struct {
	key       string
	redisPool *redis.Pool
}

func (sh *ListPublisherRedigo) Publish(data []byte) error {
	var conn redis.Conn
	if conn = sh.redisPool.Get(); conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()

	if _, err := conn.Do("RPUSH", sh.key, data); err != nil {
		return err
	}

	return nil
}

func NewListPublisherRedigo(
	key string,
	redisPool *redis.Pool,
) *ListPublisherRedigo {
	return &ListPublisherRedigo{
		key:       key,
		redisPool: redisPool,
	}
}
