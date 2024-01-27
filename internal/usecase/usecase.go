package usecase

import (
	"cirkel/user/internal/domain/dto"
	"cirkel/user/internal/repository"
	"context"
)

type usecaseInstance struct {
	cache repository.Cache
	sql   repository.Psql
}

type Usecase interface {
	// ListGenre get all lists of genres
	ListGenre(ctx context.Context) ([]*dto.ResponseGenre, error)

	// ListSkills get all lists of skills
	ListSkills(ctx context.Context) ([]*dto.ResponseSkill, error)

	// ListInstrumental get all of instruments
	ListInstrumental(ctx context.Context) ([]*dto.ResponseInstrumental, error)
}

// New creates an contract of package Usecase
func New(c repository.Cache, s repository.Psql) *usecaseInstance {
	return &usecaseInstance{
		cache: c, sql: s,
	}
}
