package service

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/integration"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/bolt"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type CardService interface {
	FindAll(ctx context.Context) ([]entity.CachedCard, error)
	Activate(ctx context.Context, card entity.CachedCard) error
	MakeTransaction(ctx context.Context, opr entity.MakeOperationRequest) error
}

type cardSrv struct {
	repo       repository.CardRepositoryBolt
	configRepo repository.ConfigRepository
	ws         integration.WSHandler
}

func NewCardService(db *bbolt.DB) CardService {
	r := bolt.NewCardRepoBolt(db)
	configRepo := bolt.NewConfigRepoBolt(db)

	ws := integration.NewWSHandler()
	return &cardSrv{
		repo:       r,
		ws:         ws,
		configRepo: configRepo,
	}
}

// FindAll implements CardService
func (c *cardSrv) FindAll(ctx context.Context) ([]entity.CachedCard, error) {
	data, err := c.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get card list, error: %w", err)
	}
	return data, nil
}

// GetCardByApplicationId implements CardService
func (c *cardSrv) Activate(ctx context.Context, card entity.CachedCard) error {

	var req = entity.SoapEnvelope{
		Soap:   "http://www.w3.org/2003/05/soap-envelope",
		Ins:    "http://bpc.ru/sv/instantissueWS/",
		Header: entity.SoapHeader{},
		Body: entity.SoapBody{
			Request: entity.GetCardInfoRequest{
				ApplId:        card.ApplicationId,
				IncludeLimits: "0",
				Lang:          "LANGENG",
			},
		},
	}

	resp, err := c.callSoap(ctx, utils.INSTANT_ISSUE, req)
	if err != nil {
		return fmt.Errorf(">> serivce: getting card information error >> %w", err)
	}

	info := resp.Body.GetCardInfoResponse.CardInfo
	log.Printf(">> change card state from [%s]", info.CardState)

	if info.CardState == "CSTE0100" {
		req = entity.SoapEnvelope{
			Soap:   "http://www.w3.org/2003/05/soap-envelope",
			Ins:    "http://bpc.ru/sv/instantissueWS/",
			Iss:    "http://bpc.ru/sv/SVXP/cardSecure",
			Header: entity.SoapHeader{},
			Body: entity.SoapBody{
				Request: entity.ProcessIBGDataRequest{
					CardID:              info.CardID,
					CardNumber:          info.CardNumber,
					CardSequentalNumber: info.SequentialNumber,
					State:               "CSTE0200",
					CardSecurity: entity.CardSecurity{
						PVV: "000",
					},
				},
			},
		}

		resp, err = c.callSoap(ctx, utils.INSTANT_ISSUE, req)
		if err != nil {
			return fmt.Errorf(">> serivce: changing card state error >> %w", err)
		}

		stateResult := resp.Body.ProcessIBGDataResponse.Text
		log.Printf(">> changed card state to [CSTE0200] operation status [%s]", stateResult)
	}

	log.Printf(">> change card status from [%s]", info.CardStatus)
	req = entity.SoapEnvelope{
		Soap:   "http://www.w3.org/2003/05/soap-envelope",
		Ins:    "http://bpc.ru/SVXP/clearing/ws",
		Iss:    "http://bpc.ru/sv/SVXP/clearing",
		Header: entity.SoapHeader{},
		Body: entity.SoapBody{
			Request: entity.OperationRequest{
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
			},
		},
	}

	_, err = c.callSoap(ctx, utils.CLEARING_WS, req)
	if err != nil {
		return fmt.Errorf(">> serivce: activating card error >> %w", err)
	}

	log.Printf("Card number [%v] is activated.", info.CardMask)

	card.CardState = "CSTE0200"
	err = c.repo.Save(ctx, card)
	if err != nil {
		return fmt.Errorf(">> serivce: activating card error >> %w", err)
	}
	return nil
}

func (c *cardSrv) callSoap(ctx context.Context, webservice string, msg entity.SoapEnvelope) (*entity.SoapReponse, error) {
	out, err := xml.MarshalIndent(&msg, " ", "  ")
	if err != nil {
		return nil, fmt.Errorf("service: marshal xml error >> %w", err)
	}

	reqBody := string(out)
	log.Printf("request data:\n%s\n", reqBody)
	cfg, _ := c.configRepo.FindByName(ctx, utils.WebserviceAddress)

	// sua lai doan nay
	strAddr := strings.TrimSuffix(cfg.Value, "/") + "/" + webservice
	log.Printf("calling webservice at [%s]", strAddr)

	resp, err := c.ws.Call(ctx, strAddr, bytes.NewBufferString(reqBody))
	if err != nil {
		return nil, fmt.Errorf("service: calling webservice error >> %w", err)
	}

	if (resp.Body.Fault != nil && resp.Body.Fault != &entity.Fault{}) {
		return nil, fmt.Errorf("service: parsing response error >> %s", resp.Body.Fault)
	}

	return resp, nil
}

// MakeTransaction implements CardService
func (c *cardSrv) MakeTransaction(ctx context.Context, opr entity.MakeOperationRequest) error {
	card := opr.Card
	var req = entity.SoapEnvelope{
		Soap:   "http://www.w3.org/2003/05/soap-envelope",
		Ins:    "http://bpc.ru/sv/instantissueWS/",
		Header: entity.SoapHeader{},
		Body: entity.SoapBody{
			Request: entity.GetCardInfoRequest{
				ApplId:        card.ApplicationId,
				IncludeLimits: "0",
				Lang:          "LANGENG",
			},
		},
	}

	resp, err := c.callSoap(ctx, utils.INSTANT_ISSUE, req)
	if err != nil {
		return fmt.Errorf(">> serivce: getting card information error >> %w", err)
	}

	info := resp.Body.GetCardInfoResponse.CardInfo
	log.Printf(">> change card state from [%s]", info.CardState)

	if info.CardState == utils.CARD_ISSUED && opr.OperationType == utils.CARD_ISSUED {
		req = entity.SoapEnvelope{
			Soap:   "http://www.w3.org/2003/05/soap-envelope",
			Ins:    "http://bpc.ru/sv/instantissueWS/",
			Iss:    "http://bpc.ru/sv/SVXP/cardSecure",
			Header: entity.SoapHeader{},
			Body: entity.SoapBody{
				Request: entity.ProcessIBGDataRequest{
					CardID:              info.CardID,
					CardNumber:          info.CardNumber,
					CardSequentalNumber: info.SequentialNumber,
					State:               utils.CARD_ACTIVATED,
					CardSecurity: entity.CardSecurity{
						PVV: utils.DEFAULT_PVV,
					},
				},
			},
		}

		resp, err = c.callSoap(ctx, utils.INSTANT_ISSUE, req)
		if err != nil {
			return fmt.Errorf(">> serivce: changing card state error >> %w", err)
		}

		stateResult := resp.Body.ProcessIBGDataResponse.Text
		log.Printf(">> changed card state to [CSTE0200] operation status [%s]", stateResult)
	}

	log.Printf(">> make operation [%s]", opr.OperationType)

	newOpr := entity.Operation{
		OperType: opr.OperationType,
		MsgType:  "MSGTPRES",
		SttlType: "STTT0000",
		OperDate: time.Now().Format(time.RFC3339)[:19],
		HostDate: time.Now().Format(time.RFC3339)[:19],
		OperAmount: &entity.OperAmount{
			AmountValue: opr.Amount,
			Currency:    "704",
		},
		Issuer: &entity.Issuer{
			ClientIDType:  "CITPCARD",
			ClientIDValue: info.CardNumber,
			CardNumber:    info.CardNumber,
			CardID:        info.CardID,
			AccountNumber: info.AccountInfo.AccountNumber,
		},
	}

	if opr.OperationType == utils.CHANGE_CARD_STATUS {
		newOpr.OperReason = "CSTS0000"
	}

	if opr.OperationType == utils.PAYMENT {
		newOpr.Issuer.ClientIDType = "CITPACCT"
		newOpr.Issuer.ClientIDValue = info.AccountInfo.AccountNumber
	}

	req = entity.SoapEnvelope{
		Soap:   "http://www.w3.org/2003/05/soap-envelope",
		Ins:    "http://bpc.ru/SVXP/clearing/ws",
		Iss:    "http://bpc.ru/sv/SVXP/clearing",
		Header: entity.SoapHeader{},
		Body: entity.SoapBody{
			Request: entity.OperationRequest{
				Operation: newOpr,
			},
		},
	}

	_, err = c.callSoap(ctx, utils.CLEARING_WS, req)
	if err != nil {
		return fmt.Errorf(">> serivce: activating card error >> %w", err)
	}

	if opr.OperationType == utils.CHANGE_CARD_STATE {

		log.Printf("Card number [%v] is activated.", info.CardMask)

		card.CardState = "CSTE0200"

		err = c.repo.Save(ctx, *card)

		if err != nil {
			return fmt.Errorf(">> serivce: activating card error >> %w", err)
		}
	}

	return nil
}
