package usecase

import (
	"cirkel/user/internal/domain/dto"
	"context"

	"github.com/cirkel-mc/goutils/errs"
	"github.com/cirkel-mc/goutils/tracer"
)

func (u *usecaseInstance) ListSkills(ctx context.Context) ([]*dto.ResponseSkill, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "Usecase:ListSkills")
	defer trace.Finish()

	results, err := u.sql.GetSkills(ctx)
	if err != nil {
		trace.SetError(err)

		return nil, errs.NewErrorWithCodeErr(err, errs.DatabaseError)
	}

	resp := make([]*dto.ResponseSkill, 0)
	for _, result := range results {
		resp = append(resp, &dto.ResponseSkill{
			Id:   result.Id,
			Name: result.Name,
			Icon: result.Image.ValueOrZero(),
		})
	}

	return resp, nil
}
