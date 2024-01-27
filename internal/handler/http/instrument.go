package http

import (
	"github.com/cirkel-mc/goutils/service"
	"github.com/cirkel-mc/goutils/tracer"
	"github.com/gofiber/fiber/v2"
)

func (h *httpInstance) getInstruments(c *fiber.Ctx) error {
	svc := service.New(c, service.User)
	ctx := c.UserContext()

	trace, ctx := tracer.StartTraceWithContext(ctx, "HttpHandler:GetInstruments")
	defer trace.Finish()

	resp, err := h.usecase.ListInstrumental(ctx)
	if err != nil {
		trace.SetError(err)

		return svc.Error(ctx, err)
	}

	return svc.OK(ctx, resp)
}
