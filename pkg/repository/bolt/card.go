package bolt

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"

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
			cc := entity.CachedCard{}
			err := json.Unmarshal(v, &cc)
			if err != nil {
				return err
			}
			if cc.CardState == "" {
				cc.CardState = "CSTE0100"
			}
			m = append(m, cc)
			return nil
		})

		return nil
	})

	if err != nil {
		return m, fmt.Errorf("get all settings error: %w", err)
	}
	sort.SliceStable(m, func(i, j int) bool {
		return m[i].CardID > m[j].CardID
	})
	return m, nil
}

// FindById implements repository.CardRepositoryBolt
func (c *cardRepoBolt) FindById(ctx context.Context, cardId int64) (*entity.CachedCard, error) {
	card := &entity.CachedCard{}
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(utils.CardTable))
		v := b.Get(itob(cardId))
		err := json.Unmarshal(v, &card)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return card, nil
}

// Save implements repository.CardRepositoryBolt
func (c *cardRepoBolt) Save(ctx context.Context, card entity.CachedCard) error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(utils.CardTable))
		if err != nil {
			return err
		}

		jo, err := json.Marshal(card)
		if err != nil {
			return err
		}
		// log.Printf("saving card [%d][%s][%s]\n", card.CardID, card.CardNumber, card.ApplicationId)
		return bucket.Put(itob(card.CardID), []byte(jo))
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
