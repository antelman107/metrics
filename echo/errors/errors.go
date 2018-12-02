// Package errors provide implementations of custom errors for the echo framework.
package errors

import (
	"fmt"
	"net/http"
	"strings"
)

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Code    int
	Cause   error
	Message string
}

// New creates a new HTTPError instance.
func New(code int, message ...string) *HTTPError {
	var e = &HTTPError{
		Code:    code,
		Message: http.StatusText(code),
	}

	if len(message) > 0 {
		e.Message = strings.Join(message, " ")
	}

	return e
}

// Wrap creates a new HTTPError instance with wrap origin error.
func Wrap(code int, err error, message ...string) *HTTPError {
	var e = &HTTPError{
		Code:    code,
		Cause:   err,
		Message: http.StatusText(code),
	}

	if len(message) > 0 {
		e.Message = strings.Join(message, " ")
	}

	return e
}

// Error makes it compatible with `error` interface.
func (e *HTTPError) Error() string {
	var msg = e.Message
	if e.Cause != nil {
		msg = e.Message + ": " + e.Cause.Error()
	}

	return fmt.Sprintf("code=%d, message=%s", e.Code, msg)
}
