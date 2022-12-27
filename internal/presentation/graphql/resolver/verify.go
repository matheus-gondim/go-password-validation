package resolver

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/matheus-gondim/go-password-validation/internal/domain/entities"
	"github.com/matheus-gondim/go-password-validation/internal/infra/graph"
)

func (r *queryResolver) Verify(ctx context.Context, input entities.VerifyPassword) (*entities.Verified, error) {
	log.Default()
	resp, err := r.VerifyUsecase.Exec(input)
	if err != nil {
		return nil,
			errors.New(http.StatusText(http.StatusInternalServerError))
	}
	return &resp, nil
}

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
