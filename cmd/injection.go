package cmd

import (
	"context"

	"github.com/cirkel-mc/goutils/abstract"
	"github.com/cirkel-mc/goutils/config"
)

func injectDependencies(cfg config.Config) (deps abstract.Dependency) {
	cfg.Injections(func(ctx context.Context) []abstract.Closer {
		master, slave := sqlDatabase()
		redis := redisDatabase(cfg.GetServiceName())

		deps = abstract.New(
			abstract.SetSQLDatabase(abstract.Master, master),
			abstract.SetSQLDatabase(abstract.Slave, slave),
			abstract.SetRedisDatabase(redis),
		)

		return []abstract.Closer{
			master, slave, redis,
		}
	})

	return
}
