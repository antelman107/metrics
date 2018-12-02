package controller

import (
	"github.com/antelman107/metrics/cmd/app/controller"
	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/definition/echo"
)

const DefControllerStats = "controller.stats"

type Stats = *controller.Stats

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefControllerStats,
			Tags: []container.Tag{
				{
					Name: echo.DefHTTPControllerTag,
				},
			},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				return controller.NewStats(), nil
			},
		})

		return nil
	})
}
