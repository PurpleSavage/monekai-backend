package authusecases

import (
	"strings"

	authentities "github.com/purplesvage/moneka-ai/internal/auth/domain/entities"
	authports "github.com/purplesvage/moneka-ai/internal/auth/domain/ports"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
)

type FindUserByEmailUseCase struct {
	repo authports.UserPersistencePort 
}

func NewFindUserByEmailUseCase(repo authports.UserPersistencePort) *FindUserByEmailUseCase {
	return &FindUserByEmailUseCase{repo: repo}
}

func (uc *FindUserByEmailUseCase) Execute(email string) (*authentities.UserEntity, error) {
	// Limpieza básica de datos
	cleanEmail := strings.TrimSpace(strings.ToLower(email))

	if cleanEmail == "" {
		return nil, domainerrors.NewAppError(400, "Invalid Email", "El correo electrónico es requerido", nil)
	}
	user, err := uc.repo.FindUserByEmail(cleanEmail)
	if err != nil {
		return nil, err
	}

	return user, nil
}