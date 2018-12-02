// Package echo provide implementations of custom functionality for the echo framework.
package echo

import (
	"github.com/labstack/echo"
)

type (
	// Echo type alias of echo.Echo
	Echo = echo.Echo

	// Group type alias of echo.Group
	Group = echo.Group

	// Context type alias of echo.Context
	Context = echo.Context

	// MiddlewareFunc type alias of echo.MiddlewareFunc.
	MiddlewareFunc = echo.MiddlewareFunc
)

var (
	// New func alias of echo.New.
	New = echo.New
)
