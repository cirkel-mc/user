package psql

import (
	"cirkel/user/internal/domain/model"
	"context"

	"github.com/cirkel-mc/goutils/logger"
	"github.com/cirkel-mc/goutils/tracer"
)

const querySkill = `
select
	id, name, image
from "user"."skills"
`

func (p *psqlRepository) GetSkills(ctx context.Context) ([]*model.Skill, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:GetSkills")
	defer trace.Finish()

	results := make([]*model.Skill, 0)
	err := p.slave.Select(ctx, &results, querySkill)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to getSkills: %s", err)

		return nil, err
	}

	return results, nil
}
