// Package echo provide implementations of custom functionality for the echo framework.
package echo

// Controller interface implementation
type Controller interface {
	Serve(e *Echo)
}
