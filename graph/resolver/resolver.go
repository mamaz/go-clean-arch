package resolver

import (
	"go-clean-arch/app/todo/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoUC usecase.TodoUC
}

func NewResolver(todoUC usecase.TodoUC) *Resolver {
	return &Resolver{
		todoUC: todoUC,
	}
}
