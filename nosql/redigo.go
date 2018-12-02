package nosql

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func getRedigoPool(address string) (*redis.Pool, error) {
	var pool = &redis.Pool{
		MaxIdle:     1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				address,
			)
		},
	}

	conn := pool.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}

	_, err := conn.Do("FLUSHDB")
	if err != nil {
		return nil, err
	}

	return pool, conn.Err()
}
