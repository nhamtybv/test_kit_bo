package bolt

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type applicationRepoBolt struct {
	db *bbolt.DB
}

func NewApplicationRepoBolt(db *bbolt.DB) repository.ApplicationConfigRepository {
	return &applicationRepoBolt{db: db}
}

// Create implements repository.ApplicationConfigRepository
func (a *applicationRepoBolt) Create(ctx context.Context) *entity.Application {
	panic("unimplemented")
}

// GetAddress implements repository.ApplicationConfigRepository
func (a *applicationRepoBolt) GetAddress(ctx context.Context) (string, error) {
	cfg, err := findConfigByName(a.db, string(utils.WebserviceAddress))
	if err != nil {
		return "", err
	}
	return cfg.Value, nil
}
