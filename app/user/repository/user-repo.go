package repository

import (
	"clean-arch/domain/user"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) GetAll(ctx context.Context) ([]user.User, error) {
	var users []user.User

	err := repo.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error on GetAll %v", err)
	}

	return users, nil
}

func (repo *UserRepo) GetByID(ID string, ctx context.Context) (user.User, error) {
	var user user.User

	err := repo.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", ID)
	if err != nil && err == sql.ErrNoRows {
		return user, nil
	}

	if err != nil {
		return user, fmt.Errorf("error on fetching user with id %s, error: %v", ID, err)
	}

	return user, nil
}

func (repo *UserRepo) Create(newuser user.UserRequest, ctx context.Context) (user.User, error) {
	createdUser := user.User{}

	// DB is set to autoincrement
	// so no need to set id before inserting
	err := repo.db.QueryRowContext(ctx, "INSERT INTO users (name, city) VALUES($1, $2) RETURNING id, name, city", newuser.City, newuser.Name).
		Scan(&createdUser.ID, &createdUser.Name, &createdUser.City)
	if err != nil {
		return user.User{}, fmt.Errorf("error on inserting user %+v, error: %v", newuser, err)
	}

	return createdUser, nil
}

func (repo *UserRepo) Delete(ID string, ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, "DELETE from users WHERE id = $1", ID)
	if err != nil {
		return fmt.Errorf("error on deleting users with id %v, error: %+v", ID, err)
	}

	return nil
}
