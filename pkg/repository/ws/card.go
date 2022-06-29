package ws

import (
	"context"
	"fmt"
	"log"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/integration"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
)

type cardRepository struct {
	config repository.ConfigRepository
	ws     integration.WSHandler
}

func NewCardRepository(config repository.ConfigRepository) repository.CardRepository {
	ws := integration.NewWSHandler()

	return &cardRepository{config: config, ws: ws}
}

// UpdateState implements repository.CardRepository
func (c *cardRepository) UpdateState(ctx context.Context, info *entity.CardInfo) error {
	url := c.config.GetConfigValue(utils.WebserviceAddress, utils.INSTANT_ISSUE)

	var req = entity.SoapEnvelope{
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

	buf, err := Otob(req)
	if err != nil {
		return fmt.Errorf(">> repo: parsing request error >> %w", err)
	}

	resp, err := c.ws.Call(ctx, url, buf)
	if err != nil {
		return fmt.Errorf("repo: calling webservice error >> %w", err)
	}

	log.Printf("update card success, result %s", resp.Body.ProcessIBGDataResponse.Text)

	return nil
}
