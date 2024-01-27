package psql

import (
	"cirkel/user/internal/domain/model"
	"context"
	"fmt"

	"github.com/cirkel-mc/goutils/logger"
	"github.com/cirkel-mc/goutils/tracer"
)

const queryClient = `
select
	id, name, client_id, client_secret, public_key, channel
from "user"."clients"
`

func (p *psqlRepository) FindClientByClientId(ctx context.Context, clientId string) (*model.Client, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:FindClientByClientId")
	defer trace.Finish()

	resp := new(model.Client)
	query := fmt.Sprintf("%s where client_id=$1", queryClient)
	err := p.slave.Get(ctx, resp, query, clientId)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to get client by client_id: %s", err)

		return nil, err
	}

	return resp, nil
}
