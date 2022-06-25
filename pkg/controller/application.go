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

type ApplicationController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type applicationCtl struct {
	s   service.ApplicationService
	ctx context.Context
}

func NewApplicationController(ctx context.Context, db *bbolt.DB) ApplicationController {

	cs := service.NewApplicationService(db)

	return &applicationCtl{
		s:   cs,
		ctx: ctx,
	}
}

// Create implements ApplicationController
func (a *applicationCtl) Create(w http.ResponseWriter, r *http.Request) {
	prd := entity.Product{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&prd); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	err := a.s.Create(a.ctx, &prd)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, prd)

}
