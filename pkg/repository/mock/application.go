package mock

import (
	"context"
	"time"

	"github.com/bxcodec/faker"
	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"go.etcd.io/bbolt"
)

type applicationRepo struct {
	db *bbolt.DB
}

func NewApplicationtRepo(db *bbolt.DB) repository.ApplicationConfigRepository {
	return &applicationRepo{db: db}
}

// Create implements repository.ApplicationConfigRepository
func (*applicationRepo) Create(ctx context.Context) *entity.Application {
	app := &entity.Application{}

	faker.FakeData(&app.Customer)
	faker.FakeData(&app.Customer.Contract.Card)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Contact.ContactData)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.SecWord)

	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Notification)
	faker.FakeData(&app.Customer.Contract.Account)
	faker.FakeData(&app.Customer.Person)
	faker.FakeData(&app.Customer.Person.PersonName)
	faker.FakeData(&app.Customer.Person.IdentityCard)
	faker.FakeData(&app.Customer.Contact.ContactData)
	faker.FakeData(&app.Customer.Address)
	faker.FakeData(&app.Customer.Address.AddressName)
	faker.FakeData(&app.Customer.Notification)

	start_date := time.Now().Format(time.RFC3339)[:10]
	app.Customer.Contract.StartDate = start_date

	return app
}
