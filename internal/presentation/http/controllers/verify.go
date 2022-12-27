package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/matheus-gondim/go-password-validation/internal/domain/entities"
	"github.com/matheus-gondim/go-password-validation/internal/domain/usecase"
)

func NewVerify(usecase usecase.IVerify) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body entities.VerifyPassword
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		resp, err := usecase.Exec(body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}
