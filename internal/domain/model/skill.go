package model

import "github.com/cirkel-mc/goutils/null"

type Skill struct {
	Id    int         `db:"id"`
	Name  string      `db:"name"`
	Image null.String `db:"image"`
}
