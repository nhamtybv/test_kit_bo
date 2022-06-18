package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/service"
	"github.com/nhamtybv/test_kit_bo/pkg/utils"
	"go.etcd.io/bbolt"
)

type ConfigController interface {
	Save(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByName(w http.ResponseWriter, r *http.Request)
}

type configCtl struct {
	cs  service.ConfigService
	ctx context.Context
}

func NewConfigController(ctx context.Context, db *bbolt.DB) ConfigController {
	cs := service.NewConfigService(db)

	return &configCtl{
		cs:  cs,
		ctx: ctx,
	}
}

// FindAll implements ConfigController
func (c *configCtl) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.cs.FindAll(c.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

// FindByName implements ConfigController
func (c *configCtl) FindByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	data, err := c.cs.FindByName(c.ctx, name)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

// Save implements ConfigController
func (c *configCtl) Save(w http.ResponseWriter, r *http.Request) {
	cfg := new(entity.Config)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cfg); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := c.cs.Save(c.ctx, cfg); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, cfg)
}
