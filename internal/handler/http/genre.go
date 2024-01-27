package http

import (
	"github.com/cirkel-mc/goutils/service"
	"github.com/cirkel-mc/goutils/tracer"
	"github.com/gofiber/fiber/v2"
)

func (h *httpInstance) getGenre(c *fiber.Ctx) error {
	svc := service.New(c, service.Auth)
	ctx := c.UserContext()

	trace, ctx := tracer.StartTraceWithContext(ctx, "HttpHandler:GetGenre")
	defer trace.Finish()

	resp, err := h.usecase.ListGenre(ctx)
	if err != nil {
		trace.SetError(err)

		return svc.Error(ctx, err)
	}

	return svc.OK(ctx, resp)
}
