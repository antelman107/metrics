package nosql

import (
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type LockRedigo struct {
	sync.RWMutex

	value     string
	redisPool *redis.Pool
}

func (l *LockRedigo) One(key string, keyDuration time.Duration) (bool, error) {
	l.Lock()
	defer l.Unlock()

	value, err := getRandString()
	if err != nil {
		return false, err
	}

	var conn redis.Conn
	if conn = l.redisPool.Get(); conn.Err() != nil {
		return false, conn.Err()
	}
	defer conn.Close()

	_, err = redis.String(conn.Do("SET", key, value, "PX", keyDuration.Nanoseconds()/1e6, "NX"))
	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	l.value = value
	return true, nil
}

func (l *LockRedigo) Iterative(key string, keyDuration time.Duration, waitDuration time.Duration) error {
	l.Lock()
	defer l.Unlock()

	var conn redis.Conn
	if conn = l.redisPool.Get(); conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()

	value, err := getRandString()
	if err != nil {
		return err
	}

	for {
		reply, err := redis.String(conn.Do("SET", key, value, "PX", keyDuration.Nanoseconds()/1e6, "NX"))
		if err != nil && err != redis.ErrNil {
			return err
		} else {
			if reply != "OK" {
				if waitDuration > 0 {
					time.Sleep(waitDuration)
				}
				continue
			}

			l.value = value
			return nil
		}
	}
}

func (l *LockRedigo) Release(key string) (bool, error) {
	var conn redis.Conn
	if conn = l.redisPool.Get(); conn.Err() != nil {
		return false, conn.Err()
	}
	defer conn.Close()

	return redis.Bool(redis.NewScript(1, `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`).Do(conn, key, l.value))
}

func NewLockRedigo(redisPool *redis.Pool) *LockRedigo {
	return &LockRedigo{
		redisPool: redisPool,
	}
}
