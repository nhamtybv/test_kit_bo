package ws

import (
	"bytes"
	"context"
	"fmt"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/integration"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
)

type applicationRepository struct {
	config repository.ConfigRepository
	ws     integration.WSHandler
}

func NewApplicationRepository(config repository.ConfigRepository) repository.ApplicationRepository {
	ws := integration.NewWSHandler()

	return &applicationRepository{config: config, ws: ws}
}

// GetCardByApplicationId implements repository.ApplicationRepository
func (a *applicationRepository) GetCardByApplicationId(ctx context.Context, applId string) (*entity.CardInfo, error) {
	url := a.config.GetConfigValue(utils.WebserviceAddress, utils.CLEARING_WS)

	var req = entity.SoapEnvelope{
		Soap:   "http://www.w3.org/2003/05/soap-envelope",
		Ins:    "http://bpc.ru/sv/instantissueWS/",
		Header: entity.SoapHeader{},
		Body: entity.SoapBody{
			Request: entity.GetCardInfoRequest{
				ApplId:        applId,
				IncludeLimits: "0",
				Lang:          "LANGENG",
			},
		},
	}

	buf, err := Otob(req)
	if err != nil {
		return nil, fmt.Errorf(">> repo: parsing request error >> %w", err)
	}

	resp, err := a.ws.Call(ctx, url, buf)
	if err != nil {
		return nil, fmt.Errorf("repo: calling webservice error >> %w", err)
	}

	return &resp.Body.GetCardInfoResponse.CardInfo, nil
}

// Mock implements repository.ApplicationRepository
func (*applicationRepository) Mock(ctx context.Context, req *entity.CardRequest) *entity.Application {
	panic("unimplemented")
}

// Create implements repository.ApplicationRepository
func (a *applicationRepository) Create(ctx context.Context, req *bytes.Buffer) (*entity.Application, error) {
	url := a.config.GetConfigValue(utils.WebserviceAddress, utils.APPLICATION_SERVICE)

	resp, err := a.ws.Call(ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("repo: calling webservice error >> %w", err)
	}

	return resp.Body.Application, nil
}
