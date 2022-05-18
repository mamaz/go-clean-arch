package server

import (
	"go-clean-arch/infrastructure/config"
	"go-clean-arch/infrastructure/logger"
	"go-clean-arch/web/api"
	"go-clean-arch/web/middlewares"

	"github.com/go-chi/chi/v5"
)

func Start() {
	r := chi.NewRouter()

	// setup config
	// if no .env file, then ENV var will be used
	config.SetWithFilePath(".", ".env")
	logger.SetupWithDebug(config.DEBUG)

	// print all env vars to make sure
	// all are set, handy for debugging
	PrintEnvVariables()

	middlewares.Set(r)
	api.Set(r)

	startServerWithGracefulShutdown(r)
}
