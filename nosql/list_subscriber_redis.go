package nosql

import (
	"github.com/go-redis/redis"
)

type ListSubscriberRedis struct {
	key         string
	redisClient redis.Cmdable
}

func (sh *ListSubscriberRedis) Get() ([]byte, error) {
	data, err := sh.redisClient.RPop(sh.key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return data, err
}

func NewListSubscriberRedis(
	key string,
	redisClient redis.Cmdable,
) *ListSubscriberRedis {
	return &ListSubscriberRedis{
		key:         key,
		redisClient: redisClient,
	}
}
