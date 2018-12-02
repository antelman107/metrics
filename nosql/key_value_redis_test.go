package nosql

import (
	"testing"
	"time"
)

func TestKeyValueRedis(t *testing.T) {
	pool, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	key := "123"
	kv := NewKeyValueRedis(pool)

	// get/set/exists
	exists, err := kv.IsExist(key)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("Exists on nonexistent")
	}

	err = kv.Set(key, time.Hour, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}

	exists, err = kv.IsExist(key)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("Not xxists on existent")
	}

	err = kv.Del(key)
	if err != nil {
		t.Fatal(err)
	}

	exists, err = kv.IsExist(key)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("Exists on nonexistent")
	}
}

func TestKeyValueRedis_Expire(t *testing.T) {
	pool, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	key := "123"
	kv := NewKeyValueRedis(pool)

	err = kv.Set(key, 500*time.Millisecond, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}

	exists, err := kv.IsExist(key)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("Not exists on existent")
	}

	time.Sleep(600 * time.Millisecond)

	exists, err = kv.IsExist(key)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("Exists on nonexistent")
	}
}
