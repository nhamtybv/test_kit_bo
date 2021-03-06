package repository

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type CardRepositoryBolt interface {
	FindAll(ctx context.Context) ([]entity.CachedCard, error)
	FindById(ctx context.Context, cardId int64) (*entity.CachedCard, error)
	CleanAll(ctx context.Context) error
	Save(ctx context.Context, card entity.CachedCard) error
}

type CardRepository interface {
	UpdateState(ctx context.Context, info *entity.CardInfo) error
}
