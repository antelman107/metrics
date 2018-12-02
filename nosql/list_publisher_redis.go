package nosql

import (
	"github.com/go-redis/redis"
)

type ListPublisherRedis struct {
	key         string
	redisClient redis.Cmdable
}

func (sh *ListPublisherRedis) Publish(data []byte) error {
	return sh.redisClient.LPush(sh.key, data).Err()
}

func NewListPublisherRedis(
	key string,
	redisClient redis.Cmdable,
) *ListPublisherRedis {
	return &ListPublisherRedis{
		key:         key,
		redisClient: redisClient,
	}
}
