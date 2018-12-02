package nosql

import (
	"sync"
	"testing"
	"time"
)

func TestOneSimple(t *testing.T) {
	duration := time.Duration(2) * time.Second

	l := NewLockMock()
	res, err := l.One("test", duration)
	if err != nil {
		t.Error("First lock error", err)
	}

	if !res {
		t.Error("First lock false")
	}

	res, err = l.One("test", duration)
	if err != nil {
		t.Error("Second lock error", err)
	}

	if res {
		t.Error("Second lock true")
	}

	<-time.NewTimer(duration).C

	res, err = l.One("test", duration)
	if err != nil {
		t.Error("Third lock error", err)
	}

	if !res {
		t.Error("Third lock false")
	}
}

func TestOneParallel(t *testing.T) {
	duration := time.Duration(2) * time.Second
	results := make([]bool, 0)

	l := NewLockMock()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		res, err := l.One("test", duration)
		if err != nil {
			t.Error("First lock error", err)
		}

		results = append(results, res)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		res, err := l.One("test", duration)
		if err != nil {
			t.Error("Second lock error", err)
		}

		results = append(results, res)
	}()

	wg.Wait()

	if len(results) != 2 {
		t.Error("Results less 2")
	}
	trues := 0
	falses := 0

	for _, val := range results {
		if val {
			trues++
		} else {
			falses++
		}
	}

	if trues != 1 {
		t.Error("Trues not 1")
	}

	if falses != 1 {
		t.Error("Falses not 1")
	}
}

func TestIterative(t *testing.T) {
	keyDuration := time.Duration(1) * time.Second
	waitDuration := time.Duration(1) * time.Millisecond
	results := make([]bool, 0)

	l := NewLockMock()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := l.Iterative("test", keyDuration, waitDuration)
		if err != nil {
			t.Error("First lock error", err)
		}

		results = append(results, true)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := l.Iterative("test", keyDuration, waitDuration)
		if err != nil {
			t.Error("Second lock error", err)
		}

		results = append(results, true)
	}()

	wg.Wait()

	if len(results) != 2 {
		t.Error("Results less 2")
	}

	for _, val := range results {
		if !val {
			t.Error("Not true result")
		}
	}
}

func TestRelease(t *testing.T) {
	l := NewLockMock()

	res, err := l.One("test", time.Minute)
	if err != nil {
		t.Error("first one error", err)
	}

	if !res {
		t.Error("first one res not true")
	}

	<-time.NewTimer(time.Duration(1) * time.Second).C
	res, err = l.One("test", time.Minute)
	if err != nil {
		t.Error("second one error", err)
	}

	if res {
		t.Error("second one res not false")
	}

	err = l.Release("test")
	if err != nil {
		t.Error("release error", err)
	}

	res, err = l.One("test", time.Minute)
	if err != nil {
		t.Error("third one error", err)
	}

	if !res {
		t.Error("third one res not true")
	}
}
