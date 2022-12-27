package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/matheus-gondim/go-password-validation/configs"
	"github.com/matheus-gondim/go-password-validation/internal/infra/graph"
	"github.com/matheus-gondim/go-password-validation/internal/presentation/graphql/resolver"
	"github.com/matheus-gondim/go-password-validation/internal/usecase"
)

func main() {
	env := configs.LoadEnv()
	c := graph.Config{Resolvers: &resolver.Resolver{VerifyUsecase: usecase.Verify{}}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.GRAPHQL_PORT)
	log.Fatal(http.ListenAndServe(":"+env.GRAPHQL_PORT, nil))
}
