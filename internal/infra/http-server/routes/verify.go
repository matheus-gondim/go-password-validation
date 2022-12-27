package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/matheus-gondim/go-password-validation/internal/domain/usecase"
	"github.com/matheus-gondim/go-password-validation/internal/presentation/http/controllers"
)

func VerifyRoutesGenerate(usecase usecase.IVerify) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/verify", func(r chi.Router) {
		r.Post("/", controllers.NewVerify(usecase))
	})

	return r
}
