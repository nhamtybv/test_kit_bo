package server

import (
	"context"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nhamtybv/test_kit_bo/pkg/controller"
	"github.com/nhamtybv/test_kit_bo/pkg/database"
	"github.com/nhamtybv/test_kit_bo/static"
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
	cardController := controller.NewCardController(ctx, bdb)

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Config
	r.HandleFunc("/api/settings", configController.FindAll).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/settings/{name}", configController.FindByName).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/settings", configController.Save).Methods(http.MethodPost, http.MethodOptions)

	// Product
	r.HandleFunc("/api/products", productController.Syns).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/products", productController.FindAll).Methods(http.MethodGet, http.MethodOptions)

	// Agents
	r.HandleFunc("/api/agents", productController.FindAllAgents).Methods(http.MethodGet, http.MethodOptions)

	// Card
	r.HandleFunc("/api/cards", cardController.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/api/applications/{application_id}", cardController.GetCardByApplicationId).Methods(http.MethodGet)

	// Application
	r.HandleFunc("/api/applications", applicationController.Create).Methods(http.MethodPost, http.MethodOptions)

	fsys := fs.FS(static.FS)
	html, _ := fs.Sub(fsys, "ui")

	r.PathPrefix("/").Handler(http.FileServer(http.FS(html)))

	r.Use(LoggerRequest)

	return r
}
