package api

import (
	"github.com/go-chi/chi/v5"
)

func Set(r *chi.Mux) {
	SetUserAPI(r)
	// place other APIs in here
}
