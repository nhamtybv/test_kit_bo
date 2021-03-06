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

type configBoltRepository struct {
	db *bbolt.DB
}

func NewConfigBoltRepository(db *bbolt.DB) repository.ConfigRepository {
	_ = db.Update(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(utils.SettingTable))

		return nil
	})
	return &configBoltRepository{db: db}
}

func (w *configBoltRepository) Save(ctx context.Context, c *entity.Config) error {
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

func (w *configBoltRepository) FindByName(ctx context.Context, name string) (*entity.Config, error) {
	return findConfigByName(w.db, name)
}

func (w *configBoltRepository) GetConfigValue(name, prefix, suffix string) string {
	val, err := findConfigByName(w.db, name)
	if err != nil {
		log.Printf("> repo get config err %s", err.Error())
		return ""
	}

	return fmt.Sprintf("%s%s%s", prefix, val.Value, suffix)
}

func (w *configBoltRepository) FindAll(ctx context.Context) (*entity.ConfigList, error) {
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
	err := db.View(func(tx *bbolt.Tx) error {
		val := tx.Bucket([]byte(utils.SettingTable)).Get([]byte(name))
		if val == nil {
			return fmt.Errorf(utils.ErrorNoDataFound)
		}
		return json.Unmarshal(val, &c)
	})
	if err != nil {
		return nil, fmt.Errorf("repo error when load config [%w]", err)
	}
	return c, nil
}
