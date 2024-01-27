package model

import (
	"time"

	"github.com/cirkel-mc/goutils/null"
)

type Genre struct {
	Id        int         `db:"id"`
	Name      string      `db:"name"`
	Image     null.String `db:"image"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt null.Time   `db:"updated_at"`
	DeletedAt null.Time   `db:"deleted_at"`
}
