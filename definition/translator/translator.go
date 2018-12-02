// Package translator provide dependency injection definitions.
package translator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"

	"github.com/antelman107/metrics/container"
)

const (
	// DefTranslator definition name.
	DefTranslator = "translator"

	// DefDefaultTranslator definition name.
	DefDefaultTranslator = "default_translator"
)

type (
	// Translator type alias of *ut.UniversalTranslator.
	Translator = *ut.UniversalTranslator

	// DefaultTranslator type alias of ut.Translator.
	DefaultTranslator = ut.Translator
)

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		var (
			locale     = en.New()
			translator = ut.New(locale)
		)

		builder.AddDefinition(container.Definition{
			Name: DefTranslator,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				return translator, nil
			},
		})

		builder.AddDefinition(container.Definition{
			Name: DefDefaultTranslator,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				return translator.GetFallback(), nil
			},
		})

		return nil
	})
}
