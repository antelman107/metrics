package nosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSubscriberRedigo(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	key := "key"
	data := "data"
	subscriber := NewListSubscriberRedigo(key, pool)

	conn := pool.Get()
	if conn.Err() != nil {
		t.Fatal(conn.Err())
	}

	for i := 1; i < 100; i++ {
		_, err := conn.Do("LPUSH", key, data)
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 1; i < 100; i++ {
		val, err := subscriber.Get()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, val, []byte("data"))
	}
}
