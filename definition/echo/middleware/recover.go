// Package middleware provide dependency injection definitions.
package middleware

import (
	"github.com/labstack/echo/middleware"

	"github.com/antelman107/metrics/container"
)

// DefEchoMiddlewareLogger recover middleware def name.
const DefEchoMiddlewareRecover = "http.middleware.recover"

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefEchoMiddlewareRecover,
			Tags: []container.Tag{{
				Name: DefEchoMiddlewareTag,
			}},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				return middleware.Recover(), nil
			},
		})

		return nil
	})
}
