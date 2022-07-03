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

type OperationService interface {
	Create(ctx context.Context, opr entity.MakeOperationRequest) error
}

type operationService struct {
	oper_ws repository.OperationRepository
	app     repository.ApplicationRepository
}

func NewOperationService(db *bbolt.DB) OperationService {
	c := bolt.NewConfigBoltRepository(db)
	op := ws.NewOperationRepository(c)
	a := ws.NewApplicationRepository(c)
	return &operationService{oper_ws: op, app: a}
}

// Create implements OperationService
func (o *operationService) Create(ctx context.Context, opr entity.MakeOperationRequest) error {

	card := opr.Card
	info, err := o.app.GetCardByApplicationId(ctx, card.ApplicationId)
	if err != nil {
		return fmt.Errorf(">> serivce: getting card information %w", err)
	}

	req := entity.OperationRequest{
		Operation: entity.Operation{
			OperType: opr.OperationType,
			MsgType:  "MSGTAUTH",
			SttlType: "STTT0010",
			OperDate: time.Now().Format(time.RFC3339)[:19],
			HostDate: time.Now().Format(time.RFC3339)[:19],
			OperAmount: &entity.OperAmount{
				AmountValue: opr.Amount,
				Currency:    "704",
			},
			Issuer: &entity.Issuer{
				ClientIDType:  "CITPACCT",
				ClientIDValue: info.AccountInfo.AccountNumber,
				CardNumber:    info.CardNumber,
				CardID:        info.CardID,
				AccountNumber: info.AccountInfo.AccountNumber,
			},
		},
	}

	if opr.OperationType == utils.PURCHASE {
		req.Operation.Issuer.ClientIDType = "CITPCARD"
		req.Operation.Issuer.ClientIDValue = info.CardNumber
	}

	oResp, err := o.oper_ws.Create(ctx, req)
	if err != nil {
		return fmt.Errorf(">> serivce: creating operation %w", err)
	}
	op := oResp.Body.OperationResponse.Operation
	log.Printf("operation id [%s] is processed with status [%s]", op.OperID, op.Status)
	return nil
}
