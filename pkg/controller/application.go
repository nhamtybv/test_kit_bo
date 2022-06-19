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
	cs  service.ProductService
	ctx context.Context
}

func NewApplicationController(ctx context.Context, db *bbolt.DB) ApplicationController {
	cs := service.NewProductService(db)

	return &applicationCtl{
		cs:  cs,
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

	utils.RespondWithJSON(w, http.StatusCreated, prd)

}
