// Package service provide implementations of various services.
package service

import "errors"

// Service an service interface.
type Service interface {
	Start(concurrency int) error
	Stop() error
	Done() <-chan struct{}
}

// ErrNotStarted triggered when service stopped before started.
var ErrNotStarted = errors.New("not started")

// ErrAlreadyStarted triggered when service started twice.
var ErrAlreadyStarted = errors.New("already started")
