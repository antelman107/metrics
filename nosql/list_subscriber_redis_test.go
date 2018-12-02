package nosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSubscriberRedis(t *testing.T) {
	conn, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	key := "key"
	data := []byte("data")
	subscriber := NewListSubscriberRedis(key, conn)

	for i := 1; i < 100; i++ {
		count, err := conn.LPush(key, data).Result()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, count, int64(i))
	}

	for i := 1; i < 100; i++ {
		val, err := subscriber.Get()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, val, []byte("data"))
	}
}
