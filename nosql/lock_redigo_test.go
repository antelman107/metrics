package nosql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLockRedigo_One(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	lock := NewLockRedigo(pool)

	key := "key"
	result, err := lock.One(key, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, result)

	result, err = lock.One(key, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, result)

	result, err = lock.Release(key)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, result)

	// Expiring
	result, err = lock.One(key, time.Millisecond*100)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, result)

	time.Sleep(time.Millisecond * 101)
	result, err = lock.Release(key)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, result)
}

func TestLockRedigo_Iterative(t *testing.T) {
	pool, err := getRedigoPool("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	lock := NewLockRedigo(pool)

	key := "key"
	err = lock.Iterative(key, time.Second*2, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}

	err = lock.Iterative(key, time.Second*2, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	//assert.Equal(t, false, result)
	//
	//result, err = lock.Release(key)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//assert.Equal(t, true, result)
}
