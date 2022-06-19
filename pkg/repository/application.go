package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type ApplicationConfigRepository interface {
	Create(ctx context.Context) *entity.Application
}
