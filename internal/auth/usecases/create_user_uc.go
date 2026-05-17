package authusecases

import (

	authentities "github.com/purplesvage/moneka-ai/internal/auth/domain/entities"
	"github.com/purplesvage/moneka-ai/internal/auth/domain/ports"
	authvalueobjects "github.com/purplesvage/moneka-ai/internal/auth/domain/valueobjects"
)

type RegisterUseCase struct{

	repo authports.UserPersistencePort
}
func NewRegisterUseCase(
	repo authports.UserPersistencePort,
) *RegisterUseCase{
	return  &RegisterUseCase{
		repo: repo,
	}
}
func (r*RegisterUseCase) Execute(createUser authvalueobjects.CreateUserVO )(*authentities.UserEntity,error){
	userCreated, err := r.repo.CreateUser(createUser)
    if err != nil {
            return nil, err
    }
	return  userCreated,nil
	
}
