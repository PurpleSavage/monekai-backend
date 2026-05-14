package auth

import (
	authHttp "github.com/purplesvage/moneka-ai/internal/auth/in/http"
	authadapters "github.com/purplesvage/moneka-ai/internal/auth/out/adpaters"
	authusecases "github.com/purplesvage/moneka-ai/internal/auth/usecases"
	sharedadapters "github.com/purplesvage/moneka-ai/internal/shared/out/adapters"
	"github.com/purplesvage/moneka-ai/internal/shared/privatemiddlewares"
	"gorm.io/gorm"
)

func Bootstrap(db *gorm.DB) *authHttp.AuthHandler {
	// 1. Adaptadores (Infraestructura)
	repo := authadapters.NewUserRepository(db)
	jwt := sharedadapters.NewJwtAdapterService()
	authProviderService := authadapters.NewAuthProviderAdapter()


	// middlewares 
	authMiddleware := privatemiddlewares.NewAuthMiddleware(
		jwt,
	)




	// 2. Casos de Uso de Apoyo (Capa de Aplicación)
	// Estos son dependencias de los casos de uso principales
	findUserByEmailUC := authusecases.NewFindUserByEmailUseCase(repo)
	updateSessionUC := authusecases.NewUpdateSessionUseCase(repo)
	
	// El RegisterUseCase suele necesitar el repo para crear el usuario
	registerUC := authusecases.NewRegisterUseCase(repo)

	// 3. Caso de Uso Principal (Orquestador)
	// Inyectamos todas las dependencias necesarias
	loginUC := authusecases.NewLoginUseCase(
		jwt, 
		repo, 
		authProviderService, 
		registerUC, 
		findUserByEmailUC, 
		updateSessionUC,
	)

	// 4. Handler (Entrada)
	return authHttp.NewAuthHandler(loginUC,authMiddleware)
}