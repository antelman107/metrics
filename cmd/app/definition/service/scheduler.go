package service

import (
	"github.com/antelman107/metrics/cmd/app/service"
	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/cmd/app/definition/cache"
	"github.com/antelman107/metrics/cmd/app/definition/nosql/queue"
	"github.com/antelman107/metrics/definition/logger"
)

const DefServiceScheduler = "service.scheduler"

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefServiceScheduler,
			Tags: []container.Tag{{
				Name: DefServiceTag,
				Args: map[string]string{
					"name": "scheduler",
				},
			}},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var log logger.Logger
				if err = ctx.Fill(logger.DefLogger, &log); err != nil {
					return nil, err
				}

				var siteRepo cache.SiteRepository
				if err = ctx.Fill(cache.DefSiteRepository, &siteRepo); err != nil {
					return nil, err
				}

				var publisherFactory queue.PublisherFactory
				if err = ctx.Fill(queue.DefNosqlQueuePublisherFactory, &publisherFactory); err != nil {
					return nil, err
				}

				return service.NewScheduler(
					siteRepo,
					publisherFactory("queue-schedule"),
					log,
				), nil
			},
		})

		return nil
	})
}
