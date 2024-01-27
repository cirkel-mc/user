package model

import (
	"time"

	"github.com/cirkel-mc/goutils/null"
)

type Client struct {
	Id           int       `db:"id"`
	Name         string    `db:"name"`
	ClientId     string    `db:"client_id"`
	ClientSecret string    `db:"client_secret"`
	PublicKey    string    `db:"public_key"`
	Channel      string    `db:"channel"`
	CreatedAt    time.Time `db:"created_at"`
	CreatedBy    null.Int  `db:"created_by"`
	UpdatedAt    null.Time `db:"updated_at"`
	UpdatedBy    null.Int  `db:"updated_by"`
	DeletedAt    null.Time `db:"deleted_at"`
	DeletedBy    null.Int  `db:"deleted_by"`
}
