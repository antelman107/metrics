// Package config provide dependency injection definitions.
package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"

	"github.com/antelman107/metrics/container"
)

// DefConfig definition name.
const DefConfig = "config"

// Config type alias of *viper.Viper.
type Config = *viper.Viper

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		var ok bool
		if _, ok = params["config_path"]; !ok {
			return errors.New("can't get required parameter config path")
		}

		var path string
		if path, ok = params["config_path"].(string); !ok {
			return errors.New(`parameter "config_path" should be string`)
		}

		builder.AddDefinition(container.Definition{
			Name: DefConfig,
			Build: func(ctx container.Context) (interface{}, error) {
				var cfg = viper.New()

				cfg.AutomaticEnv()
				cfg.SetEnvPrefix("ENV")
				cfg.SetEnvKeyReplacer(
					strings.NewReplacer(".", "_"),
				)
				cfg.SetConfigFile(path)
				cfg.SetConfigType("json")

				return cfg, cfg.ReadInConfig()
			},
		})

		return nil
	})
}
