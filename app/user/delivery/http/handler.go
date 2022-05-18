package http

import (
	"go-clean-arch/app/user/usecase"
	"go-clean-arch/domain/common/response"
	"go-clean-arch/domain/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
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
		log.Info().Msgf("error on fetching users: %+v", err)
		render.Render(w, r, response.ErrInternalServer())
		return
	}

	// WHY WE SHOULD DO THIS IN 2022?
	renderers := []render.Renderer{}
	for _, user := range users {
		renderers = append(renderers, user)
	}

	render.Render(w, r, response.SuccessResponseList(renderers))
	if err != nil {
		render.Render(w, r, response.ErrInternalServer())
	}
}

func (handler *UserHandler) getUserById(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	userFetched, err := handler.usecase.GetUserByID(userID)
	if err != nil {
		log.Info().Msgf("error getting user by id %v: %v\n", userID, err)
		render.Render(w, r, response.ErrInternalServer())
		return
	}

	err = render.Render(w, r, response.SuccessResponse(userFetched))
	if err != nil {
		render.Render(w, r, response.ErrInternalServer())
	}
}

func (handler *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	request := &user.UserRequest{}

	err := render.Bind(r, request)
	if err != nil {
		log.Info().Msgf("error on validating %+v", err)
		render.Render(w, r, response.ErrBadRequest(err))
		return
	}

	createdUser, err := handler.usecase.CreateUser(*request)
	if err != nil {
		log.Info().Msgf("error on creating user %+v", err)
		render.Render(w, r, response.ErrInternalServer())
		return
	}

	render.Render(w, r, response.SuccessResponse(createdUser))
}

func (handler *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	err := handler.usecase.DeleteUser(userID)
	if err != nil {
		render.Render(w, r, response.ErrInternalServer())
		return
	}

	w.WriteHeader(http.StatusOK)
}
