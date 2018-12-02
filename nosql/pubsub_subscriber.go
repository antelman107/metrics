package nosql

import "errors"

var ErrSubscriberFinished = errors.New("subscriber finished")

type PubSubSubscriberInterface interface {
	Receive() ([]byte, error)
	Subscribe(string) error
	Unsubscribe() error
}
