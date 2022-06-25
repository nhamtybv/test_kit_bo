package controller

import (
	"context"
	"net/http"

	"github.com/nhamtybv/test_kit_bo/pkg/service"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type ProductController interface {
	Syns(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindAllAgents(w http.ResponseWriter, r *http.Request)
}

type productCtl struct {
	cs  service.ProductService
	ctx context.Context
}

func NewProductController(ctx context.Context, db *bbolt.DB) ProductController {
	cs := service.NewProductService(db)

	return &productCtl{
		cs:  cs,
		ctx: ctx,
	}
}

// FindAll implements ProductController
func (p *productCtl) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := p.cs.FindAll(p.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}

// Syns implements ProductController
func (p *productCtl) Syns(w http.ResponseWriter, r *http.Request) {

	err := p.cs.Syns(p.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	data, err := p.cs.FindAll(p.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}

// FindAllAgent implements ProductController
func (p *productCtl) FindAllAgents(w http.ResponseWriter, r *http.Request) {
	data, err := p.cs.FindAllAgents(p.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}
