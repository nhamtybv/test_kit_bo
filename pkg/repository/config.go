package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type ConfigRepository interface {
	Save(ctx context.Context, c *entity.Config) error
	FindByName(ctx context.Context, name string) (*entity.Config, error)
	FindAll(ctx context.Context) (*entity.ConfigList, error)
}
