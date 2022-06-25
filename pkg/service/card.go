package service

import (
	"context"
	"fmt"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/bolt"
	"go.etcd.io/bbolt"
)

type CardService interface {
	FindAll(ctx context.Context) ([]entity.CachedCard, error)
}

type cardSrv struct {
	repo repository.CardRepositoryBolt
}

func NewCardService(db *bbolt.DB) CardService {
	r := bolt.NewCardRepoBolt(db)

	return &cardSrv{
		repo: r,
	}
}

// FindAll implements CardService
func (c *cardSrv) FindAll(ctx context.Context) ([]entity.CachedCard, error) {
	data, err := c.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get card list, error: %w", err)
	}
	return data, nil
}
