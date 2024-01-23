package psql

import "github.com/cirkel-mc/goutils/config/database/dbc"

type psqlRepository struct {
	master, slave dbc.SqlDbc
}

// New creates an contract of package Psql Repository
func New(m, s dbc.SqlDbc) *psqlRepository {
	return &psqlRepository{master: m, slave: s}
}
