package service

import (
	"bytes"
	"context"
	"io/fs"
	"log"
	"strconv"
	"text/template"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/mock"
	"github.com/nhamtybv/test_kit_bo/static"
	"go.etcd.io/bbolt"
)

type ApplicationService interface {
	Create(ctx context.Context, prd *entity.CardRequest) error
}

type appSrv struct {
	repo repository.ApplicationConfigRepository
	fs   fs.FS
}

func NewApplicationService(db *bbolt.DB) ApplicationService {
	r := mock.NewApplicationtRepo(db)

	return &appSrv{
		repo: r,
		fs:   static.FS,
	}
}

// Save implements ApplicationService
func (a *appSrv) Create(ctx context.Context, req *entity.CardRequest) error {
	app := a.repo.Create(ctx)

	prd := req.Product
	app.InstitutionID = strconv.Itoa(prd.InstID)
	app.Customer.Contract.ContractType = prd.ContractType
	app.Customer.Contract.ProductID = strconv.Itoa(prd.ID)
	app.Customer.Contract.Card.CardType = strconv.Itoa(prd.CardsTypes[len(prd.CardsTypes)-1].CardTypeID)
	app.AgentID = strconv.Itoa(req.AgentId)
	app.Customer.Contract.Card.Category = req.Category
	app.Customer.Contract.Account.AccountType = prd.AccountTypes[0].AccountType

	log.Printf("Institution ID: [%s], ProductID: [%s], CardType: [%s]", app.InstitutionID, app.Customer.Contract.ProductID, app.Customer.Contract.Card.CardType)

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
			Value: strconv.Itoa(v.ServiceID),
			ServiceObject: entity.ServiceObject{
				RefID:     refId,
				StartDate: app.Customer.Contract.StartDate,
			},
		}
		app.Customer.Contract.Services = append(app.Customer.Contract.Services, c)
	}

	// load template
	appTemplate := template.Must(template.ParseFS(a.fs, "templates/flow_1001.xml"))

	doc := &bytes.Buffer{}
	err := appTemplate.Execute(doc, app)
	if err != nil {
		log.Println("Error parsing template: ", err.Error())
		return err
	}
	log.Printf("Message: \n%s", doc.String())

	return nil
}
