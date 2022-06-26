package service

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"log"
	"strconv"
	"strings"
	"text/template"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/integration"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/bolt"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/mock"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"github.com/nhamtybv/test_kit_bo/static"
	"go.etcd.io/bbolt"
)

type ApplicationService interface {
	Create(ctx context.Context, prd *entity.CardRequest) error
}

type appSrv struct {
	app repository.ApplicationConfigRepository
	cfg repository.ApplicationConfigRepository
	crd repository.CardRepositoryBolt

	ws integration.WSHandler
	fs fs.FS
}

func NewApplicationService(db *bbolt.DB) ApplicationService {
	r := mock.NewApplicationtRepo()
	rb := bolt.NewApplicationRepoBolt(db)
	c := bolt.NewCardRepoBolt(db)
	ws := integration.NewWSHandler()

	return &appSrv{
		app: r,
		cfg: rb,
		crd: c,
		ws:  ws,
		fs:  static.FS,
	}
}

// Save implements ApplicationService
func (a *appSrv) Create(ctx context.Context, req *entity.CardRequest) error {
	app := a.app.Create(ctx)

	prd := req.Product
	app.InstitutionID = strconv.Itoa(prd.InstID)
	app.Customer.Contract.ContractType = prd.ContractType
	app.Customer.Contract.ProductID = strconv.Itoa(prd.ID)
	app.Customer.Contract.Card.CardType = strconv.Itoa(prd.CardsTypes[len(prd.CardsTypes)-1].CardTypeID)
	app.AgentID = strconv.Itoa(req.AgentId)
	app.Customer.Contract.Card.Category = req.Category
	app.Customer.Contract.Account.AccountType = prd.AccountTypes[0].AccountType

	log.Printf(" - Institution ID: [%s]\n - ProductID: [%s]\n - CardType: [%s]\n", app.InstitutionID, app.Customer.Contract.ProductID, app.Customer.Contract.Card.CardType)

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

	log.Printf("Message: \n%s\n", doc.String())

	strAddr, err := a.cfg.GetAddress(ctx)
	if err != nil {
		log.Println("WARNING: webservice url wasnot setted up")
	}

	strAddr = strings.TrimSuffix(strAddr, "/") + "/" + utils.APPLICATION_SERVICE

	log.Printf("calling webservice at %s", strAddr)
	resp, err := a.ws.Call(ctx, strAddr, doc)
	if err != nil {
		return fmt.Errorf("call webservice error: %w", err)
	}

	if resp.Body.Fault != nil {
		return fmt.Errorf("call webservice error: %s", resp.Body.Fault.Faultstring)
	}

	if resp.Body.Application.ApplicationStatus == "APST0008" {
		return fmt.Errorf("service: processing application error >> %s", resp.Body.Application.ApplicationID)
	}

	cardId, _ := strconv.Atoi(resp.Body.Application.Customer.Contract.Card.CardID)
	err = a.crd.Save(ctx, entity.CachedCard{
		CardID:        cardId,
		CardNumber:    resp.Body.Application.Customer.Contract.Card.CardNumber,
		ApplicationId: resp.Body.Application.ApplicationID,
	})

	if err != nil {
		log.Print("cannot add new card to database.")
	}

	log.Printf("application id %s is processed.", resp.Body.Application.ApplicationID)
	return nil
}
