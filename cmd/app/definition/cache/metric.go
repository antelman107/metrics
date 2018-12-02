package cache

import (
	"github.com/antelman107/metrics/cmd/app/cache"
	"github.com/antelman107/metrics/cmd/app/definition/nosql/keyvalue"
	"github.com/antelman107/metrics/cmd/app/definition/postgres"
	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/cmd/app/domain"

	"github.com/antelman107/metrics/definition/logger"
)

const DefMetricRepository = "repo.cache.metric"

type MetricRepository = domain.MetricRepository

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefMetricRepository,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var kvFactory keyvalue.KeyValueFactory
				if err = ctx.Fill(keyvalue.DefNosqlKeyValueFactory, &kvFactory); err != nil {
					return nil, err
				}

				var repo postgres.MetricRepository
				if err = ctx.Fill(postgres.DefMetricRepository, &repo); err != nil {
					return nil, err
				}

				var log logger.Logger
				if err = ctx.Fill(logger.DefLogger, &log); err != nil {
					return nil, err
				}

				return cache.NewStat(kvFactory(), repo), nil
			},
		})

		return nil
	})
}
