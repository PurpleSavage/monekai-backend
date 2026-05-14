package authusecases

import authports "github.com/purplesvage/moneka-ai/internal/auth/domain/ports"

type UpdateSessionUseCase struct {
	repo authports.UserPersistencePort
}
func NewUpdateSessionUseCase(repo authports.UserPersistencePort) *UpdateSessionUseCase {
	return &UpdateSessionUseCase{
		repo: repo,
	}
}
func (uc *UpdateSessionUseCase) Execute(refreshToken string, userID string) error {
	err := uc.repo.UpdateSession(refreshToken, userID)
	return err // El repo ya devuelve el AppError configurado
}