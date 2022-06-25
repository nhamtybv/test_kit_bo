package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type ApplicationConfigRepository interface {
	Create(ctx context.Context) *entity.Application
	GetAddress(ctx context.Context) (string, error)
	SaveCard(ctx context.Context, cardInfo map[string]string) error
}
