package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/service"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type CardController interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	Activate(w http.ResponseWriter, r *http.Request)
	MakeOperation(w http.ResponseWriter, r *http.Request)
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

// GetCardByApplicationId implements CardController
func (c *cardCtl) Activate(w http.ResponseWriter, r *http.Request) {

	var crd = entity.CachedCard{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&crd); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	err := c.cs.Activate(c.ctx, crd)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusCreated, "card_activated")
	}
}

// MakeOperation implements CardController
func (c *cardCtl) MakeOperation(w http.ResponseWriter, r *http.Request) {

	var opr = entity.MakeOperationRequest{
		OperationType: utils.CHANGE_CARD_STATUS,
		Card:          &entity.CachedCard{},
		Amount:        "0",
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&opr); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	err := c.cs.MakeTransaction(c.ctx, opr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusCreated, "done")
	}
}
