package usecase

import (
	"go-clean-arch/domain/todo"
	"go-clean-arch/domain/user"
	"go-clean-arch/graph/generated"
)

type TodoUC struct{}

func NewTodoUC(repo user.IUserRepo) *TodoUC {
	return &TodoUC{}
}

func (usecase *TodoUC) GetTodosByUserID(ID int) ([]*todo.Todo, error) {
	// call repo here, hardcoded for now
	var todos = make([]*todo.Todo, 1)
	todos[0] = &todo.Todo{
		ID:   1,
		Text: "some-text",
		Done: false,
		User: &user.User{
			ID:   1,
			Name: "some-name",
			City: "some-city",
		},
	}

	return todos, nil
}

func (usecase *TodoUC) CreateTodo(newTodo generated.NewTodo) (*todo.Todo, error) {
	// call repo here, hardcoded for now
	return &todo.Todo{
		ID:   1,
		Text: newTodo.Text,
		Done: false,
		User: &user.User{
			ID:   1,
			Name: "some-name",
			City: "some-city",
		},
	}, nil
}
