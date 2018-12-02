package range_queue

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlRangeQueueSubscriberFactory = "nosql.range_queue.subscriber_factory"

type SubscriberFactory func(string, string) nosql.RangeSubscriberInterface

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlRangeQueueSubscriberFactory,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return func(key string, shardName string) nosql.RangeSubscriberInterface {
					return nosql.NewRangeSubscriberRedigo(
						key,
						redis,
					)
				}, nil
			},
		})

		return nil
	})
}
