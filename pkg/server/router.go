package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nhamtybv/test_kit_bo/pkg/controller"
	"github.com/nhamtybv/test_kit_bo/pkg/database"
)

type AppRouter struct {
	Router *mux.Router
}

func NewRouter(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	bdb, err := database.NewDBDB()

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	configController := controller.NewConfigController(ctx, bdb)

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Config
	r.HandleFunc("/api/settings", configController.FindAll).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/settings", configController.Save).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/settings/{name}", configController.FindByName).Methods(http.MethodGet, http.MethodOptions)

	return r
}
