package service

import (
	"context"
	"log"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/mock"
	"go.etcd.io/bbolt"
)

type ApplicationService interface {
	Create(ctx context.Context, prd *entity.Product) error
}

type appSrv struct {
	repo repository.ApplicationConfigRepository
}

func NewApplicationService(db *bbolt.DB) ApplicationService {
	r := mock.NewApplicationtRepo(db)

	return &appSrv{
		repo: r,
	}
}

// Save implements ApplicationService
func (a *appSrv) Create(ctx context.Context, prd *entity.Product) error {
	app := a.repo.Create(ctx)
	app.InstitutionID = string(rune(prd.InstID))
	app.Customer.Contract.ContractType = prd.ContractType
	app.Customer.Contract.ProductID = string(rune(prd.ID))
	app.Customer.Contract.Card.CardType = string(rune(prd.CardsTypes[len(prd.CardsTypes)-1].CardTypeID))
	//app.Customer.Contract.Card.Category = p.Catetory
	for _, v := range prd.Services {
		log.Printf("%v", v)
		refId := "account_1"
		if v.EntityType == "ENTTCARD" {
			refId = "card_1"
		}
		if v.EntityType == "ENTTCUST" {
			refId = "customer_1"
		}

		c := entity.Service{
			Value: string(rune(v.ServiceID)),
			ServiceObject: entity.ServiceObject{
				RefID:     refId,
				StartDate: app.Customer.Contract.StartDate,
			},
		}
		app.Customer.Contract.Service = append(app.Customer.Contract.Service, c)
	}

	return nil
}
