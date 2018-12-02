package range_queue

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlRangeQueuePublisherFactory = "nosql.range_queue.publisher_factory"

type PublisherFactory func(string, string) nosql.RangePublisherInterface

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlRangeQueuePublisherFactory,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return func(key string, shardName string) nosql.RangePublisherInterface {
					return nosql.NewRangePublisherRedigo(
						key,
						redis,
					)
				}, nil
			},
		})

		return nil
	})
}
