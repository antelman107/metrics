package nosql

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type (
	KeyValueRedigo struct {
		redisPool *redis.Pool
	}
)

func NewKeyValueRedigo(pool *redis.Pool) *KeyValueRedigo {
	return &KeyValueRedigo{
		redisPool: pool,
	}
}

func (r *KeyValueRedigo) Get(key string) ([]byte, error) {
	var conn redis.Conn
	if conn = r.redisPool.Get(); conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err == redis.ErrNil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *KeyValueRedigo) Set(key string, ttl time.Duration, data []byte) (err error) {
	var conn redis.Conn
	if conn = r.redisPool.Get(); conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()

	if ttl > 0 {
		_, err = conn.Do("SET", key, data, "PX", ttl.Nanoseconds()/1e6)
	} else {
		_, err = conn.Do("SET", key, data)
	}

	return err
}

func (r *KeyValueRedigo) Del(key string) error {
	var conn redis.Conn
	if conn = r.redisPool.Get(); conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func (r *KeyValueRedigo) IsExist(key string) (bool, error) {
	var conn redis.Conn
	if conn = r.redisPool.Get(); conn.Err() != nil {
		return false, conn.Err()
	}
	defer conn.Close()

	return redis.Bool(conn.Do("EXISTS", key))
}
