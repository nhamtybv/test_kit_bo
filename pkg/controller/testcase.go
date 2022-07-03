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

type TestCaseController interface {
	Save(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

type testCaseController struct {
	cs  service.TestCaseService
	ctx context.Context
}

func NewTestCaseController(ctx context.Context, db *bbolt.DB) TestCaseController {
	cs := service.NewTestCaseService(db)

	return &testCaseController{
		cs:  cs,
		ctx: ctx,
	}
}

// FindAll implements TestCaseController
func (t *testCaseController) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := t.cs.FindAll(t.ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusOK, data)
	}
}

// FindByName implements TestCaseController
func (t *testCaseController) FindById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Save implements TestCaseController
func (t *testCaseController) Save(w http.ResponseWriter, r *http.Request) {
	var tc = entity.TestCase{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tc); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if err := t.cs.Save(t.ctx, tc); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	} else {
		utils.RespondWithJSON(w, http.StatusCreated, tc)
	}
}
