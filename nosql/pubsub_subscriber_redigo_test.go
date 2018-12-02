package nosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
)

func TestPubsubSubscriberRedigo(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}

	key := "key"
	data := []byte("data")

	subscriber := NewPubSubSubscriberRedigo(pool, time.Second*3)
	err = subscriber.Subscribe(key)
	if err != nil {
		t.Fatal(err)
	}

	conn := pool.Get()
	if conn.Err() != nil {
		t.Fatal(err)
	}

	for i := 1; i <= 100; i++ {
		_, err := conn.Do("PUBLISH", key, data)

		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 1; i <= 100; i++ {
		val, err := subscriber.Receive()

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, val, data)
	}
}

func TestPubSubSubscriberRedigo_Unsubscribe(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}

	subscriber := NewPubSubSubscriberRedigo(pool, time.Second*3)
	err = subscriber.Subscribe("somekey")
	if err != nil {
		t.Fatal(err)
	}

	err = subscriber.Unsubscribe()
	if err != nil {
		t.Fatal(err)
	}
}
