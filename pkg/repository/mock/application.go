package mock

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/bxcodec/faker"
	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
)

type applicationRepository struct {
}

// Create implements repository.ApplicationRepository
func (*applicationRepository) Create(ctx context.Context, req *bytes.Buffer) (*entity.Application, error) {
	panic("unimplemented")
}

func NewApplicationMockRepository() repository.ApplicationRepository {
	return &applicationRepository{}
}

// Create implements repository.ApplicationRepository
func (a *applicationRepository) Mock(ctx context.Context, req *entity.CardRequest) *entity.Application {
	app := &entity.Application{}

	faker.FakeData(&app.Customer)
	faker.FakeData(&app.Customer.Contract.Card)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Contact.ContactData)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.SecWord)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Address)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Address.AddressName)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Address.House)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Address.Apartment)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Person)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Person.PersonName)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Person.IdentityCard)
	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Contact.ContactData)

	faker.FakeData(&app.Customer.Contract.Card.Cardholder.Notification)
	faker.FakeData(&app.Customer.Contract.Account)
	faker.FakeData(&app.Customer.Person)
	faker.FakeData(&app.Customer.Person.PersonName)
	faker.FakeData(&app.Customer.Person.IdentityCard)
	faker.FakeData(&app.Customer.Contact.ContactData)
	faker.FakeData(&app.Customer.Address)
	faker.FakeData(&app.Customer.Address.AddressName)
	faker.FakeData(&app.Customer.Address.House)
	faker.FakeData(&app.Customer.Address.Apartment)
	faker.FakeData(&app.Customer.Notification)

	start_date := time.Now().Format(time.RFC3339)[:10]
	app.Customer.Contract.StartDate = start_date

	prd := req.Product
	app.InstitutionID = strconv.Itoa(prd.InstID)
	app.Customer.Contract.ContractType = prd.ContractType
	app.Customer.Contract.ProductID = strconv.Itoa(prd.ID)
	app.Customer.Contract.Card.CardType = strconv.Itoa(prd.CardsTypes[len(prd.CardsTypes)-1].CardTypeID)
	app.AgentID = strconv.Itoa(req.AgentId)
	app.Customer.Contract.Card.Category = req.Category
	app.Customer.Contract.Account.AccountType = prd.AccountTypes[0].AccountType

	// if req.Action == utils.ADD_SUB_CAR {
	// 	app.Customer = ""
	// }

	for _, v := range prd.Services {
		// log.Printf("%v", v)
		refId := "account_1"
		if v.EntityType == "ENTTCARD" {
			refId = "card_1"
		}
		if v.EntityType == "ENTTCUST" {
			refId = "customer_1"
		}

		c := entity.Service{
			Value: strconv.Itoa(v.ServiceID),
			ServiceObject: entity.ServiceObject{
				RefID:     refId,
				StartDate: app.Customer.Contract.StartDate,
			},
		}

		if req.Category == utils.CARD_CATEGORY_SUB && v.EntityType != "ENTTCARD" {
			continue
		}

		app.Customer.Contract.Services = append(app.Customer.Contract.Services, c)
	}

	return app
}

// GetCardByApplicationId implements repository.ApplicationConfigRepository
func (a *applicationRepository) GetCardByApplicationId(ctx context.Context, applId string) (*entity.CardInfo, error) {
	panic("unimplemented")
}
