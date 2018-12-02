// Package echo provide dependency injection definitions.
package echo

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/echo"

	"github.com/antelman107/metrics/definition/validator"
)

// DefValidator definition name.
const DefValidator = "echo.validator"

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefValidator,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var v validator.Validator
				if err = ctx.Fill(validator.DefValidator, &v); err != nil {
					return nil, err
				}

				return echo.NewValidatorWrapper(v), nil
			},
		})

		return nil
	})
}
