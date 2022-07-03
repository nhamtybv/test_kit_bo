package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type TestCaseRepository interface {
	Save(ctx context.Context, tc entity.TestCase) error
	FindAll(ctx context.Context) ([]entity.TestCase, error)
}
