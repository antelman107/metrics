package nosql

import (
	"github.com/go-redis/redis"
)

type CounterRedis struct {
	key         string
	redisClient redis.Cmdable
}

func (rc *CounterRedis) GetValue() (int64, error) {
	return rc.redisClient.Incr(rc.key).Result()
}

func NewCounterRedis(key string, redisClient redis.Cmdable) *CounterRedis {
	return &CounterRedis{
		key:         key,
		redisClient: redisClient,
	}
}
