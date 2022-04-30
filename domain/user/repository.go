package user

import "context"

// IBaseRepo is a base repo
// TODO: need to create base repo using generics
type IBaseRepo interface {
	GetAll(ctx context.Context) ([]User, error)
	Create(usrRequest UserRequest, ctx context.Context) (User, error)
	GetByID(ID string, ctx context.Context) (User, error)
	Delete(ID string, ctx context.Context) error
}

type IUserRepo interface {
	IBaseRepo
}
