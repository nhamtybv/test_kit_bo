package main

import (
	"flag"

	"github.com/nhamtybv/test_kit_bo/pkg/server"
)

func main() {
	addr := flag.String("p", "4201", "Port to run server")
	flag.Parse()
	srv := server.NewServer()

	srv.Run(":" + *addr)
}
