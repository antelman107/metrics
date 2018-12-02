package nosql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPubSubPublisherRedis(t *testing.T) {
	conn, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	key := "key"
	data := []byte("data")
	publisher := NewPubSubPublisherRedis(conn)

	pubSub := conn.Subscribe(key)

	received := [][]byte{}
	go func() {
		for {
			msg, err := pubSub.ReceiveMessage()
			if err != nil {
				t.Fatal(err)
			}
			received = append(received, []byte(msg.Payload))
		}
	}()

	for i := 1; i <= 100; i++ {
		err := publisher.Publish(key, data)

		if err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(2000 * time.Millisecond)

	assert.Equal(t, 99, len(received))
}
