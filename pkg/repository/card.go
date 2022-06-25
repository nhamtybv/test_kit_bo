package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type CardRepositoryBolt interface {
	FindAll(ctx context.Context) ([]entity.CachedCard, error)
	CleanAll(ctx context.Context) error
}
