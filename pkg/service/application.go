package service

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"log"
	"strconv"
	"text/template"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/bolt"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/mock"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/ws"
	"github.com/nhamtybv/test_kit_bo/static"
	"go.etcd.io/bbolt"
)

type ApplicationService interface {
	Create(ctx context.Context, prd *entity.CardRequest) error
}

type applicationService struct {
	mock   repository.ApplicationRepository
	ws     repository.ApplicationRepository
	config repository.ConfigRepository
	card   repository.CardRepositoryBolt
	fs     fs.FS
}

func NewApplicationService(db *bbolt.DB) ApplicationService {
	m := mock.NewApplicationMockRepository()
	c := bolt.NewConfigBoltRepository(db)
	w := ws.NewApplicationRepository(c)
	d := bolt.NewCardRepoBolt(db)

	return &applicationService{
		mock:   m,
		ws:     w,
		config: c,
		card:   d,
		fs:     static.FS,
	}
}

// Save implements ApplicationService
func (a *applicationService) Create(ctx context.Context, req *entity.CardRequest) error {
	// TODO: load data from sv and generate to sub card request
	// if req.Action == utils.ADD_SUB_CAR {
	// 	addition_infor, err := callSoapService(ctx)
	// }

	app := a.mock.Mock(ctx, req)

	log.Printf(" - Institution ID: [%s]\n - ProductID: [%s]\n - CardType: [%s]\n", app.InstitutionID, app.Customer.Contract.ProductID, app.Customer.Contract.Card.CardType)

	// load template
	templateName := fmt.Sprintf("templates/flow_%s.xml", req.Action)
	appTemplate := template.Must(template.ParseFS(a.fs, templateName))

	doc := &bytes.Buffer{}
	err := appTemplate.Execute(doc, app)
	if err != nil {
		log.Println("Error parsing template: ", err.Error())
		return err
	}

	resp, err := a.ws.Create(ctx, doc)
	if err != nil {
		return fmt.Errorf("call webservice >>  %w", err)
	}

	if resp.ApplicationStatus == "APST0008" {
		return fmt.Errorf("service: processing application >> %s", resp.ApplicationID)
	}

	log.Printf("created card [%s][%s]", resp.Customer.Contract.Card.CardID, resp.Customer.Contract.Card.CardNumber)

	cardId, _ := strconv.ParseInt(resp.Customer.Contract.Card.CardID, 10, 64)
	err = a.card.Save(ctx, entity.CachedCard{
		CardID:        cardId,
		CardNumber:    resp.Customer.Contract.Card.CardNumber,
		ApplicationId: resp.ApplicationID,
		CardState:     "CSTE0100",
		CardStatus:    resp.Customer.Contract.Card.CardStatus,
	})

	if err != nil {
		log.Print("cannot add new card to database.")
	}

	log.Printf("application id %s is processed.", resp.ApplicationID)

	return nil
}
