package bolt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"go.etcd.io/bbolt"
)

const PRODUCT_TABLE_NAME = "products"

type productRepo struct {
	db *bbolt.DB
}

func NewProductRepo(db *bbolt.DB) repository.ProductRepository {
	return &productRepo{db: db}
}

// FindAll implements repository.ProductRepository
func (p *productRepo) FindAll(ctx context.Context) (*entity.ProductList, error) {

	prds := &entity.ProductList{}

	err := p.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(TABLE_NAME))

		data := b.Get([]byte(PRODUCT_TABLE_NAME))
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
		bucket, err := tx.CreateBucketIfNotExists([]byte(PRODUCT_TABLE_NAME))
		if err != nil {
			return fmt.Errorf("creating setting bucket: %w", err)
		}

		jo, err := json.Marshal(c)
		if err != nil {
			return fmt.Errorf("json encoding: %w", err)
		}

		return bucket.Put([]byte(PRODUCT_TABLE_NAME), jo)
	})
	return err
}
