package cmd

import (
	"cirkel/user/internal/handler/http"
	"cirkel/user/internal/usecase"

	"github.com/cirkel-mc/goutils/abstract"
	"github.com/cirkel-mc/goutils/env"
	"github.com/cirkel-mc/goutils/factory/server"
	"github.com/cirkel-mc/goutils/factory/server/rest"
)

func httpHandler(deps abstract.Dependency, uc usecase.Usecase) server.ServiceFunc {
	return server.SetRestHandler(
		http.New(deps, uc),
		rest.SetHTTPPort(env.GetInt("HTTP_PORT", 9090)),
	)
}
