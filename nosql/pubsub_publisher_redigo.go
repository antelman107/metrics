package nosql

import (
	"github.com/gomodule/redigo/redis"
)

type PubSubPublisherRedigo struct {
	redisPool *redis.Pool
}

func (p *PubSubPublisherRedigo) Publish(key string, data []byte) error {
	var conn redis.Conn
	if conn = p.redisPool.Get(); conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()

	_, err := conn.Do("PUBLISH", key, data)
	if err != nil {
		return err
	}

	return nil
}

func NewPubSubPublisherRedigo(redisPool *redis.Pool) *PubSubPublisherRedigo {
	return &PubSubPublisherRedigo{
		redisPool: redisPool,
	}
}
