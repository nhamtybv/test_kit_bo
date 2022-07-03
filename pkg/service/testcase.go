package service

import (
	"context"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/bolt"
	"go.etcd.io/bbolt"
)

type TestCaseService interface {
	Save(ctx context.Context, tc entity.TestCase) error
	FindById(ctx context.Context, id int64) (*entity.TestCase, error)
	FindAll(ctx context.Context) ([]entity.TestCase, error)
}

type testCaseService struct {
	repo repository.TestCaseRepository
}

func NewTestCaseService(db *bbolt.DB) TestCaseService {
	r := bolt.NewTestCaseRepository(db)

	return &testCaseService{
		repo: r,
	}
}

// FindAll implements TestCaseService
func (t *testCaseService) FindAll(ctx context.Context) ([]entity.TestCase, error) {
	return t.repo.FindAll(ctx)
}

// FindById implements TestCaseService
func (*testCaseService) FindById(ctx context.Context, id int64) (*entity.TestCase, error) {
	panic("unimplemented")
}

// Save implements TestCaseService
func (t *testCaseService) Save(ctx context.Context, tc entity.TestCase) error {
	return t.repo.Save(ctx, tc)
}
