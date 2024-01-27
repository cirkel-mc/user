package usecase

import (
	"cirkel/user/internal/domain/dto"
	"context"

	"github.com/cirkel-mc/goutils/errs"
	"github.com/cirkel-mc/goutils/tracer"
)

func (u *usecaseInstance) ListGenre(ctx context.Context) ([]*dto.ResponseGenre, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "Usecase:ListGenre")
	defer trace.Finish()

	genres, err := u.sql.GetAllGenre(ctx)
	if err != nil {
		trace.SetError(err)

		return nil, errs.NewErrorWithCodeErr(err, errs.DatabaseError)
	}

	resp := make([]*dto.ResponseGenre, 0)
	for _, genre := range genres {
		resp = append(resp, &dto.ResponseGenre{
			Id:   genre.Id,
			Name: genre.Name,
			Icon: genre.Image.ValueOrZero(),
		})
	}

	return resp, nil
}
