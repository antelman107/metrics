// Package middleware provide dependency injection definitions.
package middleware

import (
	"github.com/antelman107/metrics/container"
	"github.com/labstack/echo/middleware"
)

// DefEchoMiddlewareLogger recover middleware def name.
const DefEchoMiddlewareLogger = "http.middleware.logger"

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefEchoMiddlewareLogger,
			Tags: []container.Tag{{
				Name: DefEchoMiddlewareTag,
			}},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				return middleware.Logger(), nil
			},
		})

		return nil
	})
}
