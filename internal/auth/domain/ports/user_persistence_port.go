package authports

import (
	authentities "github.com/purplesvage/moneka-ai/internal/auth/domain/entities"
	authvalueobjects "github.com/purplesvage/moneka-ai/internal/auth/domain/valueobjects"
)

type UserPersistencePort interface {
	CreateUser(data authvalueobjects.CreateUserVO)(*authentities.UserEntity,error)
	FindUserByEmail(email string)(*authentities.UserEntity,error)
	UpdateSession(token string, userId string) error
}