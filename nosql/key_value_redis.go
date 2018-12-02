package nosql

import (
	"time"

	"github.com/go-redis/redis"
)

type (
	KeyValueRedis struct {
		redisClient redis.Cmdable
	}
)

func NewKeyValueRedis(pool redis.Cmdable) *KeyValueRedis {
	return &KeyValueRedis{
		redisClient: pool,
	}
}

func (r *KeyValueRedis) Get(key string) ([]byte, error) {
	data, err := r.redisClient.Get(key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *KeyValueRedis) Set(key string, ttl time.Duration, data []byte) error {
	return r.redisClient.Set(key, data, ttl).Err()
}

func (r *KeyValueRedis) Del(key string) error {
	return r.redisClient.Del(key).Err()
}

func (r *KeyValueRedis) IsExist(key string) (bool, error) {
	res, err := r.redisClient.Exists(key).Result()
	if err != nil {
		return false, err
	}

	if res > 0 {
		return true, nil
	}
	return false, nil
}
