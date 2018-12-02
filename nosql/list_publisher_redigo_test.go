package nosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPublisherRedis(t *testing.T) {
	client, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	key := "key"
	data := []byte("456")
	list := NewListPublisherRedis(key, client)

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

	for i := 1; i < 100; i++ {
		d, err := client.LPop(key).Bytes()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, d, data)
	}
}
