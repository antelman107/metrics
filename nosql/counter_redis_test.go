package nosql

import (
	"testing"

	"sync"

	"github.com/stretchr/testify/assert"
)

func TestCounterRedis(t *testing.T) {
	client, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	counter := NewCounterRedis("testcounter", client)

	for i := 1; i < 100; i++ {
		val, err := counter.GetValue()

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, val, int64(i))
	}
}

func TestCounterRedis_Parallel(t *testing.T) {
	client, err := getRedisConnection("localhost:6379", 1000)
	if err != nil {
		t.Fatal(err)
	}
	counter := NewCounterRedis("testcounter", client)

	var wg sync.WaitGroup
	for i := 1; i <= 9; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			_, err := counter.GetValue()
			if err != nil {
				t.Fatal(err)
			}
		}(&wg)
	}

	wg.Wait()

	val, err := counter.GetValue()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, val, int64(10))
}
