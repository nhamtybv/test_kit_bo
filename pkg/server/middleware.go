package server

import (
	"log"
	"net/http"
)

func LoggerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%8s:%s\n", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
