// Package validator provide dependency injection definitions.
package validator

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/en"

	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/definition/translator"
)

const (
	// DefValidator definition name.
	DefValidator = "validator"

	// DefValidationTag custom validation tag name.
	DefValidationTag = "validation"
)

type (
	// Validate type func.
	Validate = func(validator Validator) error

	// Validator type alias of *validator.Validate.
	Validator = *validator.Validate

	// FieldLevel type alias of validator.FieldLevel.
	FieldLevel = validator.FieldLevel

	// FieldError type alias of validator.FieldError.
	FieldError = validator.FieldError
)

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefValidator,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var translate translator.DefaultTranslator
				if err = ctx.Fill(translator.DefDefaultTranslator, &translate); err != nil {
					return nil, err
				}

				var validate = validator.New()
				validate.RegisterTagNameFunc(func(field reflect.StructField) string {
					var (
						name      string
						tagValues = []string{
							field.Tag.Get("json"),
							field.Tag.Get("form"),
							field.Tag.Get("query"),
						}
					)

					for _, tagValue := range tagValues {
						if len(tagValue) == 0 {
							continue
						}

						name = strings.SplitN(tagValue, ",", 2)[0]
						if name == "-" {
							continue
						}

						return name
					}

					return ""
				})

				en.RegisterDefaultTranslations(validate, translate)

				err = container.Iterate(ctx, DefValidationTag, func(ctx container.Context, tag *container.Tag, name string) (err error) {
					var fn Validate
					if err = ctx.Fill(name, &fn); err != nil {
						return err
					}

					if err = fn(validate); err != nil {
						return
					}

					return nil
				})

				if err != nil {
					return nil, err
				}

				return validate, nil
			},
		})

		return nil
	})
}
