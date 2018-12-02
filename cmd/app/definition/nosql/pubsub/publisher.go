package pubsub

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlPubSubPublisherFactory = "nosql.pubsub.publisher.factory"

type Publisher = nosql.PubSubPublisherInterface
type PublisherFactory = func(string) nosql.PubSubPublisherInterface

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlPubSubPublisherFactory,
			Build: func(ctx container.Context) (_ interface{}, err error) {

				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return func(shardName string) Publisher {
					return nosql.NewPubSubPublisherRedigo(
						redis,
					)
				}, nil

			},
		})

		return nil
	})
}
