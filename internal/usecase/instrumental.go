package usecase

import (
	"cirkel/user/internal/domain/dto"
	"context"

	"github.com/cirkel-mc/goutils/errs"
	"github.com/cirkel-mc/goutils/tracer"
)

func (u *usecaseInstance) ListInstrumental(ctx context.Context) ([]*dto.ResponseInstrumental, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "Usecase:ListInstrumental")
	defer trace.Finish()

	results, err := u.sql.GetInstrumentals(ctx)
	if err != nil {
		trace.SetError(err)

		return nil, errs.NewErrorWithCodeErr(err, errs.DatabaseError)
	}

	resp := make([]*dto.ResponseInstrumental, 0)
	for _, result := range results {
		resp = append(resp, &dto.ResponseInstrumental{
			Id:   result.Id,
			Name: result.Name,
			Icon: result.Image.ValueOrZero(),
		})
	}

	return resp, nil
}
