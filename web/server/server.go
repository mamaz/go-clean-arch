package server

import (
	"go-clean-arch/infrastructure/configuration"
	"go-clean-arch/web/api"
	"go-clean-arch/web/middlewares"

	"github.com/go-chi/chi/v5"
)

func Start() {
	r := chi.NewRouter()

	// setup config
	// if no .env file, then ENV var will be used
	configuration.SetWithFilePath(".", ".env")
	middlewares.Set(r)
	api.Set(r)

	startServerWithGracefulShutdown(r)
}
