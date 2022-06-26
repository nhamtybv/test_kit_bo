package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type ProductRepository interface {
	Save(ctx context.Context, c *entity.ProductList) error
	FindByNumber(ctx context.Context, product_number string) (*entity.Product, error)
	FindAll(ctx context.Context) (*entity.ProductList, error)
	GetConnection(ctx context.Context) (string, error)
	SetConnection(str string)
	FindAgent(ctx context.Context) (*entity.Agent, error)
	SaveAgent(ctx context.Context, c *entity.Agent) error
}
