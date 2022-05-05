package usecase

import (
	"context"
	"go-clean-arch/domain/user"
)

type UserUC struct {
	repo user.IUserRepo
}

func NewUserUC(repo user.IUserRepo) *UserUC {
	return &UserUC{
		repo: repo,
	}
}

func (usecase *UserUC) GetAllUsers() ([]user.User, error) {
	return usecase.repo.GetAll(context.Background()) // TODO: use context for setting timeout
}

func (usecase *UserUC) GetUserByID(ID string) (user.User, error) {
	return usecase.repo.GetByID(ID, context.Background())
}

func (usecase *UserUC) CreateUser(newUser user.UserRequest) (user.User, error) {
	return usecase.repo.Create(newUser, context.Background())
}

func (usecase *UserUC) DeleteUser(ID string) error {
	return usecase.repo.Delete(ID, context.Background())
}
