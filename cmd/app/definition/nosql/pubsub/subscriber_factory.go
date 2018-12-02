package pubsub

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlPubSubSubscriberFactory = "nosql.pubsub.subscriber.factory"

type Subscriber = nosql.PubSubSubscriberInterface
type SubscriberFactory = func(string) nosql.PubSubSubscriberInterface

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlPubSubSubscriberFactory,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return func(shardName string) nosql.PubSubSubscriberInterface {
					return nosql.NewPubSubSubscriberRedigo(
						redis,
						0,
					)
				}, nil
			},
		})

		return nil
	})
}
