package nosql

import (
	"github.com/go-redis/redis"
)

func getRedisConnection(addr string, poolSize int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		PoolSize: poolSize,
	})
	return client, client.FlushDB().Err()
}
