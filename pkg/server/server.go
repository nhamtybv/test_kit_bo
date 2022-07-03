package server

import (
	"context"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
		Handler:      cors.Default().Handler(a.router),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// log to file
	f, err := os.OpenFile("test_kit.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	//log.SetOutput(f)

	// return right mime type
	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")

	// start server listen
	fmt.Printf("\t Server started at %s\n", addr)
	fmt.Println("\t Please change below information:")
	fmt.Println("\t  - webservice url")
	fmt.Println("\t  - oracle connection string")
	log.Fatal(ctx, srv.ListenAndServe())
}
