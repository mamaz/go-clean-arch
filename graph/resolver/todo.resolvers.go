package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-clean-arch/domain/todo"
	"go-clean-arch/graph/generated"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input generated.NewTodo) (*todo.Todo, error) {
	res, _ := r.todoUC.CreateTodo(input)

	return res, nil
}

func (r *queryResolver) Todos(ctx context.Context, userID int) ([]*todo.Todo, error) {
	res, _ := r.todoUC.GetTodosByUserID(userID)

	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
