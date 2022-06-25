package bolt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type configBolt struct {
	db *bbolt.DB
}

func NewConfigRepoBolt(db *bbolt.DB) repository.ConfigRepository {
	_ = db.Update(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(utils.SettingTable))

		return nil
	})
	return &configBolt{db: db}
}

func (w *configBolt) Save(ctx context.Context, c *entity.Config) error {
	if w.db == nil {
		return fmt.Errorf("config database is not initialed properly")
	}

	err := w.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(utils.SettingTable))
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
	return findConfigByName(w.db, name)
}

func (w *configBolt) FindAll(ctx context.Context) (*entity.ConfigList, error) {
	m := make(map[string]string)

	err := w.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(utils.SettingTable))

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

func findConfigByName(db *bbolt.DB, name string) (*entity.Config, error) {
	c := &entity.Config{}

	log.Printf("REPO: find config by: %s", name)

	err := db.View(func(tx *bbolt.Tx) error {
		log.Printf("REPO: find config by: %s", name)
		val := tx.Bucket([]byte(utils.SettingTable)).Get([]byte(name))

		if val == nil {
			return fmt.Errorf(utils.ErrorNoDataFound)
		}

		return json.Unmarshal(val, &c)
	})

	if err != nil {
		log.Printf("REPO: error when load config [%s]", err.Error())
		return nil, err

	}

	return c, nil
}
