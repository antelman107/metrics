package queue

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlQueuePublisherFactory = "nosql.queue.publisher_factory"

type PublisherFactory func(string) *nosql.ListPublisherRedigo

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlQueuePublisherFactory,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return func(key string) *nosql.ListPublisherRedigo {
					return nosql.NewListPublisherRedigo(
						key,
						redis,
					)
				}, nil
			},
		})

		return nil
	})
}
