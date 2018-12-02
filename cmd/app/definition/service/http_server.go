package service

import (
	"github.com/antelman107/metrics/cmd/app/service"
	"github.com/antelman107/metrics/container"

	_ "github.com/antelman107/metrics/cmd/app/definition/controller"
	"github.com/antelman107/metrics/definition/echo"
	"github.com/antelman107/metrics/definition/logger"
)

const DefServiceHttpServer = "service.http_server"

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefServiceHttpServer,
			Tags: []container.Tag{{
				Name: DefServiceTag,
				Args: map[string]string{
					"name": "http-server",
				},
			}},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var log logger.Logger
				if err = ctx.Fill(logger.DefLogger, &log); err != nil {
					return nil, err
				}

				var e echo.Echo
				if err = ctx.Fill(echo.DefEcho, &e); err != nil {
					return nil, err
				}

				return service.NewHttpServer(
					e,
					log,
				), nil
			},
		})

		return nil
	})
}
