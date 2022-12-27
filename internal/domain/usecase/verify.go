package usecase

import "github.com/matheus-gondim/go-password-validation/internal/domain/entities"

type IVerify interface {
	Exec(verifyPassword entities.VerifyPassword) (entities.Verified, error)
}
