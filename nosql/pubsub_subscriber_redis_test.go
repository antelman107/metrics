package nosql

import (
	"testing"

	"fmt"

	"sync"

	"github.com/stretchr/testify/assert"
)

func TestPubsubSubscriberRedis(t *testing.T) {
	conn, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}

	key := "key"
	data := []byte("data")

	subscriber := NewPubSubSubscriberRedis(conn)
	err = subscriber.Subscribe(key)
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	go func(wg *sync.WaitGroup) {
		fmt.Printf("Begin subscribe listening\n")

		for i := 1; i <= 100; i++ {
			val, err := subscriber.Receive()

			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, val, data)
		}
	}(&wg)

	for i := 1; i <= 100; i++ {
		err := conn.Publish(key, data).Err()

		if err != nil {
			t.Fatal(err)
		}
	}

	wg.Wait()
}

func TestPubSubSubscriberRedis_Unsubscribe(t *testing.T) {
	conn, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}

	subscriber := NewPubSubSubscriberRedis(conn)
	err = subscriber.Subscribe("somedata")
	if err != nil {
		t.Fatal(err)
	}

	err = subscriber.Unsubscribe()
	if err != nil {
		t.Fatal(err)
	}
}
