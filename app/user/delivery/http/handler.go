package http

import (
	"encoding/json"
	"fmt"
	"go-clean-arch/app/user/usecase"
	"go-clean-arch/domain/user"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserHandler struct {
	//repo user.IUserRepo
	usecase *usecase.UserUC
}

func NewHandler(uc *usecase.UserUC) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}

func (handler *UserHandler) SetRoute(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.getAllUsers)
		r.Post("/", handler.createUser)
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", handler.getUserById)
			r.Delete("/", handler.deleteUser)
		})
	})
}

func (handler *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.usecase.GetAllUsers()
	if err != nil {
		w.Write([]byte("error"))
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(`{"error": "500"}`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"data": %s}`, string(usersJSON))))
}

func (handler *UserHandler) getUserById(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	userFetched, err := handler.usecase.GetUserByID(userID)
	if err != nil {
		log.Println("error: ", err)
		r.Response.StatusCode = 500
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	var empty user.User

	if userFetched == empty {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}

	userJSON, err := json.Marshal(userFetched)
	if err != nil {
		w.Write([]byte(`{"error": "500"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(userJSON))
}

func (handler *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	request := &user.UserRequest{}

	err := render.Bind(r, request)
	if err != nil {
		errMessage := fmt.Sprintf(`{
			"error": "400",
			"message": %v
		}`, err.Error())
		w.Write([]byte(errMessage))
		return
	}

	createdUser, err := handler.usecase.CreateUser(*request)

	if err != nil {
		log.Printf("error on creating user %+v", err)
		w.Write([]byte(`{"error": "500"}`))
		return
	}

	userJSON, err := json.Marshal(createdUser)
	if err != nil {
		w.Write([]byte(`{"error": "500", "message": "error on marshalling user"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(userJSON))
}

func (handler *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := ""
	err := handler.usecase.DeleteUser(userID)
	if err != nil {
		w.Write([]byte("error"))
	}
}
