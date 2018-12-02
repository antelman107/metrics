package nosql

import (
	"github.com/go-redis/redis"
)

type PubSubPublisherRedis struct {
	redisClient *redis.Client
}

func (p *PubSubPublisherRedis) Publish(key string, data []byte) error {
	return p.redisClient.Publish(key, data).Err()
}

func NewPubSubPublisherRedis(redisClient *redis.Client) *PubSubPublisherRedis {
	return &PubSubPublisherRedis{
		redisClient: redisClient,
	}
}
