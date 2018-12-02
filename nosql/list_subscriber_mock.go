package nosql

import (
	"fmt"
	"sync"
)

type ListSubscriberMock struct {
	sync.Mutex

	data [][]byte
}

func (sh *ListSubscriberMock) Get() ([]byte, error) {
	sh.Lock()
	defer sh.Unlock()

	if len(sh.data) > 0 {
		res := sh.data[len(sh.data)-1]
		sh.data = sh.data[:len(sh.data)-1]
		return res, nil
	}

	return nil, fmt.Errorf("No subscriber data")
}

func NewListSubscriberMock(
	data [][]byte,
) *ListSubscriberMock {
	return &ListSubscriberMock{
		data: data,
	}
}
