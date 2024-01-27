package psql

import (
	"cirkel/user/internal/domain/model"
	"context"

	"github.com/cirkel-mc/goutils/logger"
	"github.com/cirkel-mc/goutils/tracer"
)

const queryInstrumental = `
select
	id, name, image
from "user"."instruments"
where deleted_at is null
`

func (p *psqlRepository) GetInstrumentals(ctx context.Context) ([]*model.Instrument, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:GetInstrumentals")
	defer trace.Finish()

	results := make([]*model.Instrument, 0)
	err := p.slave.Select(ctx, &results, queryInstrumental)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed get GetInstrumentals: %s", err)

		return nil, err
	}

	return results, nil
}
