// Package echo provide dependency injection definitions.
package echo

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/echo/errors"
)

// DefErrorHandler definition name.
const DefErrorHandler = "echo.error_handler"

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefErrorHandler,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				// See original echo.DefaultHTTPErrorHandler
				// https://godoc.org/github.com/labstack/echo#Echo.DefaultHTTPErrorHandler
				return func(err error, c echo.Context) {
					var (
						e      = c.Echo()
						logger = c.Logger()
						code   = http.StatusInternalServerError
						msg    = http.StatusText(code)
					)

					switch he := err.(type) {
					case *echo.HTTPError:
						code = he.Code
						msg = he.Message.(string)
					case *errors.HTTPError:
						code = he.Code
						msg = he.Message
					}

					if e.Debug {
						msg = err.Error()
					}

					logger.Error(err)

					if !c.Response().Committed {
						if c.Request().Method == echo.HEAD {
							err = c.NoContent(code)
						} else {
							err = c.JSON(code, echo.Map{"message": msg})
						}
						if err != nil {
							logger.Error(err)
						}
					}
				}, nil
			},
		})

		return nil
	})
}
