package api

import (
	"go-clean-arch/app/user/delivery/http"
	"go-clean-arch/app/user/repository"
	"go-clean-arch/app/user/usecase"
	"go-clean-arch/infrastructure/database"

	"github.com/go-chi/chi/v5"
)

func SetUserAPI(r *chi.Mux) {
	db := database.CreateDBConnection("localhost", 5432, "postgres", "password01", "postgres")
	repo := repository.NewUserRepo(db)
	usecase := usecase.NewUserUC(repo)
	usrHandler := http.NewHandler(usecase)
	usrHandler.SetRoute(r)
}
