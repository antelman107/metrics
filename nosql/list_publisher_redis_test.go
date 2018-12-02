package nosql

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestListPublisherRedigo(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	key := "key"
	data := []byte("456")
	list := NewListPublisherRedigo(key, pool)

	for i := 1; i < 100; i++ {
		err := list.Publish(data)

		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 1; i < 100; i++ {
		err := list.Publish(data)

		if err != nil {
			t.Fatal(err)
		}
	}

	conn := pool.Get()
	if conn.Err() != nil {
		t.Fatal(conn.Err())
	}

	for i := 1; i < 100; i++ {
		d, err := redis.Bytes(conn.Do("LPOP", key))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, d, data)
	}
}
