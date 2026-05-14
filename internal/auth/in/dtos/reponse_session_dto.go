package authdtos

import authentities "github.com/purplesvage/moneka-ai/internal/auth/domain/entities"

type ResponseSessionDto struct {
	UserData    authentities.UserEntity
	AccessToken string
}