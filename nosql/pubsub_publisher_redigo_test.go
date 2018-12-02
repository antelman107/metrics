package nosql

import (
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestPubSubPublisherRedigo(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	key := "key"
	data := []byte("data")
	publisher := NewPubSubPublisherRedigo(pool)

	conn := pool.Get()
	if conn.Err() != nil {
		t.Fatal(conn.Err())
	}
	pubSubConn := redis.PubSubConn{Conn: conn}
	pubSubConn.Subscribe(key)

	received := []string{}
	go func() {
		for {
			switch n := pubSubConn.Receive().(type) {
			case redis.Message:
				received = append(received, string(n.Data))
			case error:
			}
		}
	}()

	for i := 1; i < 100; i++ {
		err := publisher.Publish(key, data)

		if err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(500 * time.Millisecond)

	assert.Equal(t, len(received), 99)
}
