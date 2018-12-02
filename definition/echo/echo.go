// Package echo provide dependency injection definitions.
package echo

import (
	"errors"
	"path"

	"github.com/labstack/gommon/log"

	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/echo"

	"github.com/antelman107/metrics/definition/config"
	"github.com/antelman107/metrics/definition/echo/middleware"
)

const (
	// DefEcho definition name.
	DefEcho = "echo"

	// DefHTTPControllerTag custom http controller tag name.
	DefHTTPControllerTag = "http.controller.tag"
)

type (
	// Echo type alias of *echo.Echo
	Echo = *echo.Echo

	// Controller type alias of http.Controller
	Controller = echo.Controller
)

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		var (
			configPath string
			ok         bool
		)
		if configPath, ok = params["config_path"].(string); !ok {
			return errors.New(`parameter "config_path" should be string`)
		}

		builder.AddDefinition(container.Definition{
			Name: DefEcho,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var conf config.Config
				if err = ctx.Fill(config.DefConfig, &conf); err != nil {
					return nil, err
				}

				var e = echo.New()
				e.Debug = conf.GetBool("http.debug")

				switch conf.GetString("http.level") {
				case "debug":
					e.Logger.SetLevel(log.DEBUG)
				case "info":
					e.Logger.SetLevel(log.INFO)
				case "warn":
					e.Logger.SetLevel(log.WARN)
				case "error":
					e.Logger.SetLevel(log.ERROR)
				case "off":
					e.Logger.SetLevel(log.OFF)
				}

				var resourcePath = path.Dir(configPath)
				e.Static("/static", resourcePath+"/public")

				if err = ctx.Fill(DefValidator, &e.Validator); err != nil {
					return nil, err
				}

				if err = ctx.Fill(DefErrorHandler, &e.HTTPErrorHandler); err != nil {
					return nil, err
				}

				if err = container.Iterate(
					ctx,
					DefHTTPControllerTag,
					func(ctx container.Context, _ *container.Tag, name string) (err error) {
						var c echo.Controller
						if err = ctx.Fill(name, &c); err != nil {
							return err
						}

						c.Serve(e)

						return nil
					},
				); err != nil {
					return nil, err
				}

				if err = container.Iterate(
					ctx,
					middleware.DefEchoMiddlewareTag,
					func(ctx container.Context, _ *container.Tag, name string) (err error) {
						var mw middleware.MiddlewareFunc
						if err = ctx.Fill(name, &mw); err != nil {
							return err
						}

						e.Use(mw)

						return nil
					},
				); err != nil {
					return nil, err
				}

				return e, nil
			},
		})

		return nil
	})
}
