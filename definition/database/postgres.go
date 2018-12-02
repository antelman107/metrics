// Package postgres provide dependency injection definitions.
package database

import (
	"strings"

	"github.com/antelman107/metrics/container"
	"github.com/iqoption/nap"
	_ "github.com/lib/pq" // database driver

	"github.com/antelman107/metrics/definition/config"
)

// DefPostgres definition name.
const DefPostgres = "db.postgres"

// Postgres type alias of *nap.DB.
type Postgres = *nap.DB

// Definition init func.
func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefPostgres,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var cfg config.Config
				if err = ctx.Fill(config.DefConfig, &cfg); err != nil {
					return nil, err
				}

				var db *nap.DB
				if db, err = nap.Open(
					"postgres",
					strings.Join(cfg.GetStringSlice("postgres"), ";"),
				); err != nil {
					return nil, err
				}

				if err = db.Ping(); err != nil {
					return nil, err
				}

				return db, nil
			},
			Close: func(obj interface{}) {
				obj.(*nap.DB).Close()
			},
		})

		return nil
	})
}
