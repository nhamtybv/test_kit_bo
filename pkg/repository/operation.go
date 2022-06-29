package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type OperationRepository interface {
	Create(ctx context.Context, opr interface{}) (*entity.SoapReponse, error)
}
