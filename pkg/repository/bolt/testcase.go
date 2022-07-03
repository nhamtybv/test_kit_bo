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

type testCaseRepository struct {
	db *bbolt.DB
}

func NewTestCaseRepository(db *bbolt.DB) repository.TestCaseRepository {
	_ = db.Update(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(utils.SettingTable))

		return nil
	})
	return &testCaseRepository{db: db}
}

func (t *testCaseRepository) Save(ctx context.Context, tc entity.TestCase) error {
	err := t.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(utils.TestCaseTable))
		if err != nil {
			return fmt.Errorf("creating testcase bucket: %w", err)
		}

		if tc.Id == 0 {
			_id, _ := bucket.NextSequence()
			tc.Id = int64(_id)
		}

		jo, err := json.Marshal(tc)
		if err != nil {
			return fmt.Errorf("json encoding: %w", err)
		}

		return bucket.Put(utils.Itob(tc.Id), jo)
	})
	return err
}

// FindAll implements repository.TestCaseRepository
func (t *testCaseRepository) FindAll(ctx context.Context) ([]entity.TestCase, error) {
	result := []entity.TestCase{}

	err := t.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(utils.TestCaseTable))

		b.ForEach(func(k, v []byte) error {
			cc := entity.TestCase{}
			err := json.Unmarshal(v, &cc)
			if err != nil {
				return err
			}

			result = append(result, cc)
			return nil
		})

		return nil
	})
	if err != nil {
		return result, fmt.Errorf("loading testcases >> %w", err)
	}

	return result, nil
}
