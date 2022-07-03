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
}

type cardController struct {
	cs  service.CardService
	op  service.OperationService
	ctx context.Context
}

func NewCardController(ctx context.Context, db *bbolt.DB) CardController {
	cs := service.NewCardService(db)
	op := service.NewOperationService(db)

	return &cardController{
		cs:  cs,
		op:  op,
		ctx: ctx,
	}
}

// FindAll implements CardController
func (c *cardController) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.cs.FindAll(c.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}

// GetCardByApplicationId implements CardController
func (c *cardController) Activate(w http.ResponseWriter, r *http.Request) {

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
