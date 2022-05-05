package api

import (
	"go-clean-arch/app/user/delivery/http"
	"go-clean-arch/app/user/repository"
	"go-clean-arch/app/user/usecase"
	"go-clean-arch/infrastructure/configuration"
	"go-clean-arch/infrastructure/database"

	"github.com/go-chi/chi/v5"
)

func SetUserAPI(r *chi.Mux) {
	db := database.CreateDBConnection(
		configuration.POSTGRES_HOST,
		configuration.POSTGRES_PORT,
		configuration.POSTGRES_USERNAME,
		configuration.POSTGRES_PASSWORD,
		configuration.POSTGRES_DATABASE,
	)
	repo := repository.NewUserRepo(db)
	usecase := usecase.NewUserUC(repo)

	usrHandler := http.NewHandler(usecase)

	// route should be explicitly set
	usrHandler.SetRoute(r)
}
