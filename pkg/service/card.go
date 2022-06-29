package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/bolt"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/ws"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type CardService interface {
	FindAll(ctx context.Context) ([]entity.CachedCard, error)
	Activate(ctx context.Context, card entity.CachedCard) error
}

type cardService struct {
	card    repository.CardRepositoryBolt
	config  repository.ConfigRepository
	app     repository.ApplicationRepository
	card_ws repository.CardRepository
	oper_ws repository.OperationRepository
}

func NewCardService(db *bbolt.DB) CardService {
	r := bolt.NewCardRepoBolt(db)
	c := bolt.NewConfigBoltRepository(db)
	a := ws.NewApplicationRepository(c)
	w := ws.NewCardRepository(c)
	o := ws.NewOperationRepository(c)

	return &cardService{
		card:    r,
		config:  c,
		app:     a,
		card_ws: w,
		oper_ws: o,
	}
}

// FindAll implements CardService
func (c *cardService) FindAll(ctx context.Context) ([]entity.CachedCard, error) {
	data, err := c.card.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get card list, error: %w", err)
	}
	return data, nil
}

// GetCardByApplicationId implements CardService
func (c *cardService) Activate(ctx context.Context, card entity.CachedCard) error {

	info, err := c.app.GetCardByApplicationId(ctx, card.ApplicationId)
	if err != nil {
		return fmt.Errorf(">> serivce: getting card information error >> %w", err)
	}

	log.Printf(">> change card state from [%s]", info.CardState)

	if info.CardState == "CSTE0100" {
		err = c.card_ws.UpdateState(ctx, info)
		if err != nil {
			return fmt.Errorf(">> serivce: changing card state >> %w", err)
		}
		log.Printf(">> changed card state to [CSTE0200]")
	}

	log.Printf(">> change card status from [%s]", info.CardStatus)

	req := entity.OperationRequest{
		Operation: entity.Operation{
			OperType: utils.CHANGE_CARD_STATUS,
			MsgType:  "MSGTPRES",
			SttlType: "STTT0000",
			OperDate: time.Now().Format(time.RFC3339)[:19],
			HostDate: time.Now().Format(time.RFC3339)[:19],
			OperAmount: &entity.OperAmount{
				AmountValue: "0",
				Currency:    "704",
			},
			OperReason: "CSTS0000",
			Issuer: &entity.Issuer{
				ClientIDType:  "CITPCARD",
				ClientIDValue: info.CardNumber,
				CardNumber:    info.CardNumber,
				CardID:        info.CardID,
			},
		},
	}

	// TO DO: change response
	_, err = c.oper_ws.Create(ctx, req)

	if err != nil {
		return fmt.Errorf(">> serivce: activating card error >> %w", err)
	}

	log.Printf("Card number [%v] is activated.", info.CardMask)

	card.CardState = "CSTE0200"
	err = c.card.Save(ctx, card)
	if err != nil {
		return fmt.Errorf(">> serivce: activating card error >> %w", err)
	}

	return nil
}
