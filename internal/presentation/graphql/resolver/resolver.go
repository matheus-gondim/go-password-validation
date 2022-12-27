package resolver

import "github.com/matheus-gondim/go-password-validation/internal/domain/usecase"

type Resolver struct {
	VerifyUsecase usecase.IVerify
}
