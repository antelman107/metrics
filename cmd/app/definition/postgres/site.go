package postgres

import (
	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/cmd/app/postgres"

	"github.com/antelman107/metrics/definition/database"
	"github.com/antelman107/metrics/definition/logger"
)

const DefSiteRepository = "repo.postgres.site"

type SiteRepository = domain.SiteRepository

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefSiteRepository,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var db database.Postgres
				if err = ctx.Fill(database.DefPostgres, &db); err != nil {
					return nil, err
				}

				var log logger.Logger
				if err = ctx.Fill(logger.DefLogger, &log); err != nil {
					return nil, err
				}

				return postgres.NewSite(db, log), nil
			},
		})

		return nil
	})
}
