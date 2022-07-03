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

type OperationController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type operationController struct {
	op_repo service.OperationService
	ctx     context.Context
}

func NewOperationController(ctx context.Context, db *bbolt.DB) OperationController {

	op := service.NewOperationService(db)

	return &operationController{
		op_repo: op,
		ctx:     ctx,
	}
}

// MakeOperation implements CardController
func (o *operationController) Create(w http.ResponseWriter, r *http.Request) {

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

	err := o.op_repo.Create(o.ctx, opr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusCreated, "done")
	}
}
