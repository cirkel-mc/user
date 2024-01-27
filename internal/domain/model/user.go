package model

import (
	"cirkel/user/internal/domain/constant"
	"time"

	"github.com/cirkel-mc/goutils/null"
)

type User struct {
	Id         int                 `db:"id"`
	RoleId     int                 `db:"role_id"`
	Username   string              `db:"username"`
	Password   string              `db:"password"`
	Email      string              `db:"email"`
	Status     constant.UserStatus `db:"status"`
	IsPartner  bool                `db:"is_partner"`
	VerifiedAt null.Time           `db:"verified_at"`
	SkillId    null.Int            `db:"skill_id"`
	Preference null.String         `db:"preference"`
	CreatedAt  time.Time           `db:"created_at"`
	CreatedBy  null.Int            `db:"created_by"`
	UpdatedAt  null.Time           `db:"updated_at"`
	UpdatedBy  null.Int            `db:"updated_by"`
	DeletedAt  null.Time           `db:"deleted_at"`
	DeletedBy  null.Int            `db:"deleted_by"`

	// Foreign Key
	Role  *Role  `db:"ro"`
	Skill *Skill `db:"sk"`
}
