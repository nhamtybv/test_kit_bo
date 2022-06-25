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
	productController := controller.NewProductController(ctx, bdb)
	applicationController := controller.NewApplicationController(ctx, bdb)

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Config
	r.HandleFunc("/api/settings", configController.FindAll).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/settings", configController.Save).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/settings/{name}", configController.FindByName).Methods(http.MethodGet, http.MethodOptions)

	// Product
	r.HandleFunc("/api/products", productController.Syns).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/products", productController.FindAll).Methods(http.MethodGet, http.MethodOptions)

	// Agents
	r.HandleFunc("/api/agents", productController.FindAllAgents).Methods(http.MethodGet, http.MethodOptions)

	// Appluication
	r.HandleFunc("/api/applications", applicationController.Create).Methods(http.MethodPost, http.MethodOptions)

	return r
}
