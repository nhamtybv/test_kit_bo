package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Application interface {
	Run(addr string)
}

type appServer struct {
	router *mux.Router
}

func NewServer() Application {
	ctx := context.Background()
	r := NewRouter(ctx)

	return &appServer{
		router: r,
	}
}

// Run implements Application
func (a *appServer) Run(addr string) {
	ctx := context.Background()
	srv := &http.Server{
		Handler:      a.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// start server listen
	log.Printf("Server started at %s", addr)
	log.Fatal(ctx, srv.ListenAndServe())
}
