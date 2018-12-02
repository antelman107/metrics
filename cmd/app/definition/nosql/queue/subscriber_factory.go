package queue

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlQueueSubscriberFactory = "nosql.queue.subscriber_factory"

type SubscriberFactory func(string) *nosql.ListSubscriberRedigo

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlQueueSubscriberFactory,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return func(key string) *nosql.ListSubscriberRedigo {
					return nosql.NewListSubscriberRedigo(
						key,
						redis,
					)
				}, nil
			},
		})

		return nil
	})
}
