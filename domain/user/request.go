package user

import "net/http"

type UserRequest struct {
	Name string
	City string
}

func (a *UserRequest) Bind(r *http.Request) error {
	return nil
}
