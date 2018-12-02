package nosql

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type LockRedis struct {
	sync.RWMutex

	value       string
	redisClient redis.Cmdable
}

func (l *LockRedis) One(key string, keyDuration time.Duration) (bool, error) {
	l.Lock()
	defer l.Unlock()

	value, err := getRandString()
	if err != nil {
		return false, err
	}

	res, err := l.redisClient.SetNX(key, value, keyDuration).Result()
	if err != nil {
		return false, err
	}

	if res {
		l.value = value
	}
	return res, nil
}

func (l *LockRedis) Iterative(key string, keyDuration time.Duration, waitDuration time.Duration) error {
	l.Lock()
	defer l.Unlock()

	value, err := getRandString()
	if err != nil {
		return err
	}
	for {
		res, err := l.redisClient.SetNX(key, value, keyDuration).Result()
		if err != nil {
			return err
		} else {
			if !res {
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

func (l *LockRedis) Release(key string) (bool, error) {
	result, err := l.redisClient.Eval(`
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`, []string{key}, l.value).Result()

	if err != nil {
		return false, err
	}

	if result.(int64) == 0 {
		return false, nil
	}

	return true, nil
}

func NewLockRedis(redisClient redis.Cmdable) *LockRedis {
	return &LockRedis{
		redisClient: redisClient,
	}
}
