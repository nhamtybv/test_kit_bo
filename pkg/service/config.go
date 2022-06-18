package service

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"go.etcd.io/bbolt"
)

type ConfigService interface {
	Save(ctx context.Context, cfg *entity.Config) error
	FindByName(ctx context.Context, name string) (*entity.Config, error)
	FindAll(ctx context.Context) (*entity.ConfigList, error)
}

type configSrv struct {
	repo repository.ConfigRepository
}

func NewConfigService(db *bbolt.DB) ConfigService {
	r := repository.NewWebServiceConfigBolt(db)

	return &configSrv{
		repo: r,
	}
}

// FindAll implements ConfigService
func (c *configSrv) FindAll(ctx context.Context) (*entity.ConfigList, error) {
	data, err := c.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// FindByName implements ConfigService
func (c *configSrv) FindByName(ctx context.Context, name string) (*entity.Config, error) {
	data, err := c.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Save implements ConfigService
func (c *configSrv) Save(ctx context.Context, cfg *entity.Config) error {
	return c.repo.Save(ctx, cfg)
}
