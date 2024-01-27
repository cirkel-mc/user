package http

import (
	"cirkel/user/internal/usecase"

	"github.com/cirkel-mc/goutils/abstract"
	"github.com/cirkel-mc/goutils/validation"
	"github.com/gofiber/fiber/v2"
)

type httpInstance struct {
	middleware abstract.Middleware
	validator  validation.Validation
	usecase    usecase.Usecase
}

func New(deps abstract.Dependency, uc usecase.Usecase) abstract.RESTHandler {
	return &httpInstance{
		middleware: deps.GetMiddleware(),
		validator:  validation.New(),
		usecase:    uc,
	}
}

func (h *httpInstance) Router(r fiber.Router) {
	v1 := r.Group("/v1")
	// v1.Use(h.middleware.HTTPSignatureValidate)

	{
		genrePath := v1.Group("/genre")
		genrePath.Get("", h.getGenre)
	}

	{
		skillPath := v1.Group("/skill")
		skillPath.Get("", h.getSkill)
	}

	{
		instrumentPath := v1.Group("/instrument")
		instrumentPath.Get("", h.getInstruments)
	}
}
