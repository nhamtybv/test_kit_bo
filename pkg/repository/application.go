package repository

import (
	"bytes"
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type ApplicationRepository interface {
	Mock(ctx context.Context, req *entity.CardRequest) *entity.Application
	GetCardByApplicationId(ctx context.Context, applId string) (*entity.CardInfo, error)
	Create(ctx context.Context, req *bytes.Buffer) (*entity.Application, error)
}
