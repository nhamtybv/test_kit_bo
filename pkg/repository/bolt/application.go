package bolt

import (
	"context"
	"encoding/json"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type applicationRepoBolt struct {
	db *bbolt.DB
}

func NewApplicationRepoBolt(db *bbolt.DB) repository.ApplicationConfigRepository {
	return &applicationRepoBolt{db: db}
}

// Create implements repository.ApplicationConfigRepository
func (a *applicationRepoBolt) Create(ctx context.Context) *entity.Application {
	panic("unimplemented")
}

// GetAddress implements repository.ApplicationConfigRepository
func (a *applicationRepoBolt) GetAddress(ctx context.Context) (string, error) {
	cfg, err := findConfigByName(a.db, string(utils.WebserviceAddress))
	if err != nil {
		return "", err
	}
	return cfg.Value, nil
}

// SaveCard implements repository.ApplicationConfigRepository
func (a *applicationRepoBolt) SaveCard(ctx context.Context, cardInfo map[string]string) error {

	err := a.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(utils.CardTable))
		if err != nil {
			return err
		}

		card_number := cardInfo["card_number"]

		jo, err := json.Marshal(cardInfo)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(card_number), []byte(jo))
	})

	return err
}
