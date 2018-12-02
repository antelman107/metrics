package controller

import (
	"github.com/antelman107/metrics/cmd/app/controller"
	"github.com/antelman107/metrics/cmd/app/definition/cache"
	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/definition/echo"
)

const DefControllerMetric = "controller.metric"

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefControllerMetric,
			Tags: []container.Tag{{
				Name: echo.DefHTTPControllerTag,
			}},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var metricRepo cache.MetricRepository
				if err = ctx.Fill(cache.DefMetricRepository, &metricRepo); err != nil {
					return nil, err
				}

				var siteRepo cache.SiteRepository
				if err = ctx.Fill(cache.DefSiteRepository, &siteRepo); err != nil {
					return nil, err
				}

				return controller.NewMetric(
					metricRepo,
					siteRepo,
				)
			},
		})

		return nil
	})
}
