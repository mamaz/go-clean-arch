package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Metadata struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type AppResponse struct {
	Metadata
	Data render.Renderer `json:"data"`
}

func (response AppResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AppResponseList struct {
	Metadata
	Data []render.Renderer `json:"data"`
}

func (response AppResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
