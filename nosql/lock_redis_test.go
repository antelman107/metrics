package nosql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLockRedis_One(t *testing.T) {
	pool, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	lock := NewLockRedis(pool)

	key := "key"
	result, err := lock.One(key, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, result)

	result, err = lock.One(key, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, result)

	result, err = lock.Release(key)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, result)

	// Expiring
	result, err = lock.One(key, time.Millisecond*100)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, result)

	time.Sleep(time.Millisecond * 101)
	result, err = lock.Release(key)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, result)
}

func TestLockRedis_Iterative(t *testing.T) {
	pool, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	lock := NewLockRedis(pool)

	key := "key"
	err = lock.Iterative(key, time.Minute, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}

	result, err := lock.One(key, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, result)

	result, err = lock.Release(key)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, result)
}
