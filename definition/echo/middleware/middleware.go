// Package middleware provide dependency injection definitions.
package middleware

import "github.com/antelman107/metrics/echo"

// DefEchoMiddlewareTag custom middleware tag name.
const DefEchoMiddlewareTag = "http.middleware.tag"

// MiddlewareFunc type alias of echo.MiddlewareFunc
type MiddlewareFunc = echo.MiddlewareFunc
