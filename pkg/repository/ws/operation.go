package ws

import (
	"context"
	"fmt"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/integration"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
)

type operationRepository struct {
	config repository.ConfigRepository
	ws     integration.WSHandler
}

func NewOperationRepository(config repository.ConfigRepository) repository.OperationRepository {
	ws := integration.NewWSHandler()
	return &operationRepository{config: config, ws: ws}
}

// MakeOperation implements repository.OperationRepository
func (o *operationRepository) Create(ctx context.Context, opr interface{}) (*entity.SoapReponse, error) {
	url := o.config.GetConfigValue(utils.WebserviceAddress, utils.CLEARING_WS)

	var req = entity.SoapEnvelope{
		Soap:   "http://www.w3.org/2003/05/soap-envelope",
		Ins:    "http://bpc.ru/sv/instantissueWS/",
		Header: entity.SoapHeader{},
		Body: entity.SoapBody{
			Request: opr,
		},
	}

	buf, err := Otob(req)
	if err != nil {
		return nil, fmt.Errorf(">> repo: parsing request error >> %w", err)
	}

	resp, err := o.ws.Call(ctx, url, buf)
	if err != nil {
		return nil, fmt.Errorf("repo: calling webservice error >> %w", err)
	}

	return resp, nil
}
