package api

import (
	"clean-arch/app/user/delivery/http"
	"clean-arch/app/user/repository"
	"clean-arch/app/user/usecase"
	"clean-arch/infrastructure/database"

	"github.com/go-chi/chi/v5"
)

func SetUserAPI(r *chi.Mux) {
	db := database.CreateDBConnection("localhost", 5432, "postgres", "password01", "postgres")
	repo := repository.NewUserRepo(db)
	usecase := usecase.NewUserUC(repo)
	usrHandler := http.NewHandler(usecase)
	usrHandler.SetRoute(r)
}
