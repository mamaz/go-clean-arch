package api

import (
	"go-clean-arch/app/todo/usecase"
	"go-clean-arch/graph/generated"
	"go-clean-arch/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

func SetGraphqlAPI(r *chi.Mux) {
	todoUC := usecase.NewTodoUC(nil)

	resolver := resolver.NewResolver(*todoUC)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)
}
