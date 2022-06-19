package bolt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"go.etcd.io/bbolt"
)

const TABLE_NAME = "settings"

type configBolt struct {
	db *bbolt.DB
}

func NewConfigRepoBolt(db *bbolt.DB) repository.ConfigRepository {
	return &configBolt{db: db}
}

func (w *configBolt) Save(ctx context.Context, c *entity.Config) error {
	if w.db == nil {
		return fmt.Errorf("config database is not initialed properly")
	}

	err := w.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(TABLE_NAME))
		if err != nil {
			return fmt.Errorf("creating setting bucket: %w", err)
		}

		jo, err := json.Marshal(c)
		if err != nil {
			return fmt.Errorf("json encoding: %w", err)
		}

		return bucket.Put([]byte(c.Name), jo)
	})
	return err
}

func (w *configBolt) FindByName(ctx context.Context, name string) (*entity.Config, error) {
	c := &entity.Config{}
	err := w.db.View(func(tx *bbolt.Tx) error {
		val := tx.Bucket([]byte(TABLE_NAME)).Get([]byte(name))
		if val == nil {
			c.Name = name
			c.Value = "no_data_found"

			w.Save(ctx, c)
		}
		return json.Unmarshal(val, &c)
	})

	return c, err
}

func (w *configBolt) FindAll(ctx context.Context) (*entity.ConfigList, error) {
	m := make(map[string]string)

	err := w.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(TABLE_NAME))

		b.ForEach(func(k, v []byte) error {
			tmp := &entity.Config{}
			err := json.Unmarshal(v, &tmp)
			if err != nil {
				return err
			}
			m[string(k)] = string(tmp.Value)
			return nil
		})
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("get all settings error: %w", err)
	}

	lst := &entity.ConfigList{
		Count:   len(m),
		Configs: m,
	}

	return lst, nil
}
