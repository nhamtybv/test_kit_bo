package controller

import (
	"context"
	"encoding/json"
	"log"
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

type configController struct {
	cs  service.ConfigService
	ctx context.Context
}

func NewConfigController(ctx context.Context, db *bbolt.DB) ConfigController {
	cs := service.NewConfigService(db)

	return &configController{
		cs:  cs,
		ctx: ctx,
	}
}

// FindAll implements ConfigController
func (c *configController) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.cs.FindAll(c.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}

// FindByName implements ConfigController
func (c *configController) FindByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Printf("CONTROLLER: find config by: %s", name)
	data, err := c.cs.FindByName(c.ctx, name)

	if err != nil {
		if err.Error() == utils.ErrorNoDataFound {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}

// Save implements ConfigController
func (c *configController) Save(w http.ResponseWriter, r *http.Request) {
	cfg := new(entity.Config)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cfg); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if err := c.cs.Save(c.ctx, cfg); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusCreated, cfg)
	}
}
