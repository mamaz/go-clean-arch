package server

import (
	"go-clean-arch/web/api"
	"go-clean-arch/web/middlewares"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Start() {
	r := chi.NewRouter()

	middlewares.Set(r)
	api.Set(r)

	log.Println("Starting server in port :8080")
	http.ListenAndServe(":8080", r)
}
