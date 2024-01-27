package psql

import (
	"cirkel/user/internal/domain/model"
	"context"
	"fmt"

	"github.com/cirkel-mc/goutils/logger"
	"github.com/cirkel-mc/goutils/tracer"
)

const queryUserGenre = `
select
	us.id "us.id", us.email "us.email", us.username "us.username", us.status "us.status",
	gr.id "gr.id", gr.name "gr.name", gr.image "gr.image"
from "user"."user_genre" as ug
inner join "user"."users" as us on us.id = ug.user_id
inner join "user"."genre" as gr on gr.id = ug.genre_id
`

func (p *psqlRepository) GetUserGenreByUser(ctx context.Context, userId int) ([]*model.UserGenre, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:GetUserGenreByUser")
	defer trace.Finish()

	query := fmt.Sprintf("%s where us.id=$1", queryUserGenre)
	var resp = make([]*model.UserGenre, 0)

	err := p.slave.Select(ctx, resp, query, userId)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to get GetUserGenreByUser: %s", err)

		return nil, err
	}

	return resp, nil
}

func (p *psqlRepository) GetUserGenreByGenre(ctx context.Context, genreId int) ([]*model.UserGenre, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:GetUserGenreByGenre")
	defer trace.Finish()

	query := fmt.Sprintf("%s where gr.id=$1", queryUserGenre)
	var resp = make([]*model.UserGenre, 0)

	err := p.slave.Select(ctx, resp, query, genreId)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to get GetUserGenreByGenre: %s", err)

		return nil, err
	}

	return resp, nil
}
