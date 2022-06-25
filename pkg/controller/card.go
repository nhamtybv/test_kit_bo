package controller

import (
	"context"
	"net/http"

	"github.com/nhamtybv/test_kit_bo/pkg/service"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type CardController interface {
	FindAll(w http.ResponseWriter, r *http.Request)
}

type cardCtl struct {
	cs  service.CardService
	ctx context.Context
}

func NewCardController(ctx context.Context, db *bbolt.DB) CardController {
	cs := service.NewCardService(db)

	return &cardCtl{
		cs:  cs,
		ctx: ctx,
	}
}

// FindAll implements CardController
func (c *cardCtl) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.cs.FindAll(c.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}
