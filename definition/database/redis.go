package database

import (
	"net"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/antelman107/metrics/container"

	"github.com/antelman107/metrics/definition/config"
)

const DefRedis = "db.redis"

type Redis = *redis.Pool

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		builder.AddDefinition(container.Definition{
			Name: DefRedis,
			Build: func(ctx container.Context) (_ interface{}, err error) {
				var cfg config.Config
				if err = ctx.Fill(config.DefConfig, &cfg); err != nil {
					return nil, err
				}

				var pool = &redis.Pool{
					MaxIdle:     500,
					IdleTimeout: 240 * time.Second,
					Dial: func() (redis.Conn, error) {
						return redis.Dial(
							"tcp",
							net.JoinHostPort(
								cfg.GetString("redis.host"),
								cfg.GetString("redis.port"),
							),
						)
					},
				}

				var conn = pool.Get()
				if conn.Err() != nil {
					return nil, conn.Err()
				}
				defer conn.Close()

				if _, err = conn.Do("PING"); err != nil {
					return nil, err
				}

				return pool, nil
			},
		})

		return nil
	})
}
