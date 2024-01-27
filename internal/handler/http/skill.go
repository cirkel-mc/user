package http

import (
	"github.com/cirkel-mc/goutils/service"
	"github.com/cirkel-mc/goutils/tracer"
	"github.com/gofiber/fiber/v2"
)

func (h *httpInstance) getSkill(c *fiber.Ctx) error {
	svc := service.New(c, service.User)
	ctx := c.UserContext()

	trace, ctx := tracer.StartTraceWithContext(ctx, "HttpHandler:GetSkills")
	defer trace.Finish()

	resp, err := h.usecase.ListSkills(ctx)
	if err != nil {
		trace.SetError(err)

		return svc.Error(ctx, err)
	}

	return svc.OK(ctx, resp)
}
