package usecase

import "cirkel/user/internal/repository"

type usecaseInstance struct {
	cache repository.Cache
	sql   repository.Psql
}

type Usecase interface{}

// New creates an contract of package Usecase
func New(c repository.Cache, s repository.Psql) *usecaseInstance {
	return &usecaseInstance{
		cache: c, sql: s,
	}
}
