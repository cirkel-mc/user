package cache

import "github.com/cirkel-mc/goutils/config/database/rdc"

type cacheRepository struct {
	client rdc.Rdc
}

// New creates an contract of package Cache Repository
func New(c rdc.Rdc) *cacheRepository {
	return &cacheRepository{client: c}
}
