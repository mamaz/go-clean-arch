package api

import (
	"go-clean-arch/app/user/delivery/http"
	"go-clean-arch/app/user/repository"
	"go-clean-arch/app/user/usecase"
	"go-clean-arch/infrastructure/config/secrets"
	"go-clean-arch/infrastructure/database"

	"github.com/go-chi/chi/v5"
)

func SetUserAPI(r *chi.Mux) {
	db := database.CreateDBConnection(
		secrets.POSTGRES_HOST,
		secrets.POSTGRES_PORT,
		secrets.POSTGRES_USERNAME,
		secrets.POSTGRES_PASSWORD,
		secrets.POSTGRES_DATABASE,
	)
	repo := repository.NewUserRepo(db)
	usecase := usecase.NewUserUC(repo)

	// route should be explicitly set
	http.NewHandler(usecase).SetRoute(r)
}
