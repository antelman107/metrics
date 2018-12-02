package keyvalue

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlKeyValueFactory = "nosql.keyvalue.factory"

type (
	KeyValue        nosql.KeyValueInterface
	KeyValueFactory func() KeyValue
)

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlKeyValueFactory,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return func() KeyValue {
					return nosql.NewKeyValueRedigo(
						redis,
					)
				}, nil
			},
		})

		return nil
	})
}
