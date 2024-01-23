package cmd

import (
	"cirkel/user/internal/repository/cache"
	"cirkel/user/internal/repository/psql"
	"cirkel/user/internal/usecase"

	"github.com/cirkel-mc/goutils/abstract"
	"github.com/cirkel-mc/goutils/config"
	"github.com/cirkel-mc/goutils/factory/server"
	"github.com/cirkel-mc/goutils/middleware"
)

func Serve(cfg config.Config) *server.Server {
	deps := injectDependencies(cfg)

	// middleware
	mdl := middleware.New(nil)
	deps.SetMiddleware(mdl)

	// repositories
	psqlRepo := psql.New(deps.GetSQLDatabase(abstract.Master).Database(), deps.GetSQLDatabase(abstract.Slave).Database())
	cacheRepo := cache.New(deps.GetRedisDatabase().Client())

	// usecase
	uc := usecase.New(psqlRepo, cacheRepo)

	// initiates services
	svc := server.NewApplicationService(
		server.SetConfiguration(cfg),
		server.SetDependencies(deps),
		httpHandler(deps, uc),
	)

	// initiates server
	srv := server.New(svc)
	return srv
}
