package service

import (
	"context"
)

type OperationService interface {
	Create(ctx context.Context, opr interface{}) error
}

type operationService struct {
}

// Create implements OperationService
func (operationService) Create(ctx context.Context, opr interface{}) error {
	panic("unimplemented")
}

func NewOperationService() OperationService {
	return operationService{}
}

/*
// MakeTransaction implements CardService
func (c *cardService) MakeTransaction(ctx context.Context, opr entity.MakeOperationRequest) error {
	card := opr.Card

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

*/
