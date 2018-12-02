// Package middleware provide dependency injection definitions.
package middleware

import (
	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/cmd/app/definition/controller"
	"github.com/antelman107/metrics/definition/echo/middleware"
)

// DefEchoMiddlewareStats recover middleware def name.
const DefEchoMiddlewareStats = "http.middleware.stats"

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefEchoMiddlewareStats,
			Tags: []container.Tag{{
				Name: middleware.DefEchoMiddlewareTag,
			}},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var stats controller.Stats
				if err = ctx.Fill(controller.DefControllerStats, &stats); err != nil {
					return nil, err
				}

				return stats.Process, nil
			},
		})

		return nil
	})
}
