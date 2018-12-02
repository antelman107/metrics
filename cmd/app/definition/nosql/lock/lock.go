package lock

import (
	"github.com/antelman107/metrics/container"
	"github.com/antelman107/metrics/nosql"
	"github.com/gomodule/redigo/redis"

	"github.com/antelman107/metrics/definition/database"
)

const DefNosqlLock = "nosql.lock"

type Lock = nosql.LockInterface

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefNosqlLock,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var redis database.Redis
				if err = ctx.Fill(database.DefRedis, &redis); err != nil {
					return nil, err
				}

				return nosql.NewLockRedigo(
					redis,
				), nil
			},
		})

		return nil
	})
}
