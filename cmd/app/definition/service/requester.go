package service

import (
	"github.com/antelman107/metrics/cmd/app/definition/cache"
	"github.com/antelman107/metrics/cmd/app/service"
	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/cmd/app/definition/nosql/queue"
	"github.com/antelman107/metrics/definition/logger"
)

const DefServiceRequester = "service.requester"

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefServiceRequester,
			Tags: []container.Tag{{
				Name: DefServiceTag,
				Args: map[string]string{
					"name": "requester",
				},
			}},
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var log logger.Logger
				if err = ctx.Fill(logger.DefLogger, &log); err != nil {
					return nil, err
				}

				var metricRepo cache.MetricRepository
				if err = ctx.Fill(cache.DefMetricRepository, &metricRepo); err != nil {
					return nil, err
				}

				var subscriberFactory queue.SubscriberFactory
				if err = ctx.Fill(queue.DefNosqlQueueSubscriberFactory, &subscriberFactory); err != nil {
					return nil, err
				}

				return service.NewRequester(
					metricRepo,
					subscriberFactory("queue-schedule"),
					log,
				), nil
			},
		})

		return nil
	})
}
