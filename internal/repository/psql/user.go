package psql

import (
	"cirkel/user/internal/domain/model"
	"context"
	"fmt"

	"github.com/cirkel-mc/goutils/logger"
	"github.com/cirkel-mc/goutils/tracer"
)

const queryUser = `
select
	us.id, us.username, us.email, us.password, us.status,
	us.verified_at, us.is_partner, us.role_id, us.skill_id, us.preference,
	ro.id "ro.id", ro.name "ro.name", ro.key "ro.key", ro.type "ro.type",
	sk.id "sk.id", sk.name "sk.name", sk.image "sk.image"
from "user"."users" as us
inner join "user"."roles" as ro on ro.id = us.role_id
left join "user"."skills" as sk on sk.id = us.skill_id
where us.deleted_at is null
`

func (p *psqlRepository) FindUserById(ctx context.Context, id int) (*model.User, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:FindUserById")
	defer trace.Finish()

	query := fmt.Sprintf("%s and us.id=$1", queryUser)
	resp := &model.User{}
	resp.Role = &model.Role{}
	resp.Skill = &model.Skill{}

	err := p.slave.Get(ctx, resp, query, id)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to get findUserById: %s", err)

		return nil, err
	}

	return resp, nil
}

func (p *psqlRepository) FindUserByUsernameOrEmail(ctx context.Context, val string) (*model.User, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:FindUserByUsernameOrEmail")
	defer trace.Finish()

	query := fmt.Sprintf("%s and us.email=$1 or us.username=$1", queryUser)
	resp := &model.User{}
	resp.Role = &model.Role{}
	resp.Skill = &model.Skill{}

	err := p.slave.Get(ctx, resp, query, val)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to get findUserByUsernameOrEmail: %s", err)

		return nil, err
	}

	return resp, nil
}

func (p *psqlRepository) CreateUser(ctx context.Context, user *model.User) error {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:CreateUser")
	defer trace.Finish()

	query := `
	insert into "user"."users" (role_id, username, email, password, status, created_at, is_partner, skill_id)
	values ($1, $2, $3, $4, $5, $6, $7, $8) returning id
	`

	err := p.master.QueryRowx(
		ctx, query,
		user.Role.Id,
		user.Username,
		user.Email,
		user.Password,
		user.Status,
		user.CreatedAt,
		user.IsPartner,
		user.SkillId,
	).StructScan(user)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to createUser: %s", err)

		return err
	}

	return nil
}

func (p *psqlRepository) UpdateUser(ctx context.Context, user *model.User) error {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:UpdateUser")
	defer trace.Finish()

	query := `
	update "user"."users"
		status=$2, verified_at=$3, updated_at=$4, password=$5,
		skill_id=$6, preference=$7, is_partner=$8, updated_by=$9
	where id=$1
	`

	err := p.master.Preparex(
		ctx, query,
		user.Id, user.Status,
		user.VerifiedAt, user.UpdatedAt,
		user.Password, user.SkillId,
		user.Preference, user.IsPartner,
		user.UpdatedBy,
	)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to updateUser: %s", err)

		return err
	}

	return nil
}

func (p *psqlRepository) DeleteUser(ctx context.Context, user *model.User) error {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PsqlRepository:DeleteUser")
	defer trace.Finish()

	query := `
	update "user"."users"
		status=$2, deleted_at=$3, deleted_by=$4
	where id=$1
	`

	err := p.master.Preparex(
		ctx, query,
		user.Id, user.Status,
		user.DeletedAt, user.DeletedBy,
	)
	if err != nil {
		trace.SetError(err)
		logger.Log.Errorf(ctx, "failed to deleteUser: %s", err)

		return err
	}

	return nil
}
