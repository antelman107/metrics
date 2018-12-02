package nosql

import (
	"testing"

	"sync"

	"github.com/stretchr/testify/assert"
)

func TestCounterRedigo(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	counter := NewCounterRedigo("testcounter", pool)

	for i := 1; i < 100; i++ {
		val, err := counter.GetValue()

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, val, int64(i))
	}
}

func TestCounterRedigo_Parallel(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	counter := NewCounterRedigo("testcounter", pool)

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
