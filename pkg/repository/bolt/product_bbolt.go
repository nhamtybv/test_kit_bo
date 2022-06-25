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

type productRepo struct {
	db *bbolt.DB
}

func NewProductRepo(db *bbolt.DB) repository.ProductRepository {
	return &productRepo{db: db}
}

// FindAll implements repository.ProductRepository
func (p *productRepo) FindAll(ctx context.Context) (*entity.ProductList, error) {

	prds := &entity.ProductList{
		Count:    0,
		Products: []entity.Product{},
	}

	err := p.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(utils.ProductTable))

		data := b.Get([]byte(utils.ProductTable))
		err := json.Unmarshal(data, &prds)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("get products error: %w", err)
	}
	return prds, nil
}

// FindByNumber implements repository.ProductRepository
func (*productRepo) FindByNumber(ctx context.Context, product_number string) (*entity.Product, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Save implements repository.ProductRepository
func (p *productRepo) Save(ctx context.Context, c *entity.ProductList) error {
	if p.db == nil {
		return fmt.Errorf("config database is not initialed properly")
	}

	err := p.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(utils.ProductTable))
		if err != nil {
			return fmt.Errorf("creating setting bucket: %w", err)
		}

		jo, err := json.Marshal(c)
		if err != nil {
			return fmt.Errorf("json encoding: %w", err)
		}

		return bucket.Put([]byte(utils.ProductTable), jo)
	})
	return err
}

// GetConnection implements repository.ProductRepository
func (p *productRepo) GetConnection(ctx context.Context) (string, error) {
	c := &entity.Config{}
	err := p.db.View(func(tx *bbolt.Tx) error {
		val := tx.Bucket([]byte(utils.SettingTable)).Get([]byte(utils.OracleConnectionKey))
		if val == nil {
			c.Name = string(utils.OracleConnectionKey)
			c.Value = "no_data_found"
		}
		return json.Unmarshal(val, &c)
	})

	return c.Value, err
}

// FindAgent implements repository.ProductRepository
func (p *productRepo) FindAgent(ctx context.Context) (*entity.Agent, error) {
	agent := &entity.Agent{}

	err := p.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(utils.ProductTable))

		data := b.Get([]byte(utils.AgentTable))
		err := json.Unmarshal(data, &agent)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("get products error: %w", err)
	}
	return agent, nil
}

// SaveAgent implements repository.ProductRepository
func (p *productRepo) SaveAgent(ctx context.Context, c *entity.Agent) error {
	if p.db == nil {
		return fmt.Errorf("config database is not initialed properly")
	}

	err := p.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(utils.ProductTable))
		if err != nil {
			return fmt.Errorf("creating setting bucket: %w", err)
		}

		jo, err := json.Marshal(c)
		if err != nil {
			return fmt.Errorf("json encoding: %w", err)
		}

		return bucket.Put([]byte(utils.AgentTable), jo)
	})
	return err
}
