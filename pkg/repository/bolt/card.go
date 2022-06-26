package bolt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type cardRepoBolt struct {
	db *bbolt.DB
}

func NewCardRepoBolt(db *bbolt.DB) repository.CardRepositoryBolt {
	return &cardRepoBolt{db: db}
}

// CleanAll implements repository.CardRepositoryBolt
func (c *cardRepoBolt) CleanAll(ctx context.Context) error {
	panic("unimplemented")
}

// FindAll implements repository.CardRepositoryBolt
func (c *cardRepoBolt) FindAll(ctx context.Context) ([]entity.CachedCard, error) {
	m := []entity.CachedCard{}

	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(utils.CardTable))

		b.ForEach(func(k, v []byte) error {

			tmp := map[string]string{}

			err := json.Unmarshal(v, &tmp)
			// log.Printf("card: %v", tmp)
			if err != nil {
				// log.Printf("card: %v", tmp)
				return err
			}

			card_id := tmp[utils.CARD_ID]
			card_number := tmp[utils.CARD_NUMBER]

			m = append(m, entity.CachedCard{
				CardID:     card_id,
				CardNumber: card_number,
			})
			return nil
		})
		return nil
	})

	if err != nil {
		return m, fmt.Errorf("get all settings error: %w", err)
	}

	return m, nil
}
