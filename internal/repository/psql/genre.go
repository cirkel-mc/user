package psql

import (
	"cirkel/user/internal/domain/model"
	"context"

	"github.com/cirkel-mc/goutils/logger"
	"github.com/cirkel-mc/goutils/tracer"
)

const queryGenre = `
select
	gr.id, gr.name, gr.image
from "user"."genre" as gr
where gr.deleted_at is null
`

func (p *psqlRepository) GetAllGenre(ctx context.Context) ([]*model.Genre, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:GetAllGenre")
	defer trace.Finish()

	results := make([]*model.Genre, 0)
	err := p.slave.Select(ctx, &results, queryGenre)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to getAllGenre: %s", err)

		return nil, err
	}

	return results, nil
}
