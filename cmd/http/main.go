package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/matheus-gondim/go-password-validation/configs"
	server "github.com/matheus-gondim/go-password-validation/internal/infra/http-server"
	"github.com/matheus-gondim/go-password-validation/internal/infra/http-server/routes"
	"github.com/matheus-gondim/go-password-validation/internal/usecase"
)

func main() {
	env := configs.LoadEnv()

	server.New(env.HTTP_PORT, []*chi.Mux{
		routes.VerifyRoutesGenerate(usecase.Verify{}),
	}).Run()
}
