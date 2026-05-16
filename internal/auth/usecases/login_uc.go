package authusecases

import (
	authentities "github.com/purplesvage/moneka-ai/internal/auth/domain/entities"
	"github.com/purplesvage/moneka-ai/internal/auth/domain/ports"
	authvalueobjects "github.com/purplesvage/moneka-ai/internal/auth/domain/valueobjects"
	"github.com/purplesvage/moneka-ai/internal/shared/domain/exceptions"
	sharedports "github.com/purplesvage/moneka-ai/internal/shared/domain/ports"
)


type LoginUseCase struct{
	jwtService  sharedports.JwtPort
	repo authports.UserPersistencePort
	authService authports.AuthProviderPort
    registerUseCase *RegisterUseCase
    findUserByEmailUC  *FindUserByEmailUseCase
    updateSessionUC *UpdateSessionUseCase
}

func NewLoginUseCase(jwt sharedports.JwtPort,
	repo authports.UserPersistencePort,
	authService authports.AuthProviderPort,
    registerUC *RegisterUseCase, 
    findUserByEmailUC *FindUserByEmailUseCase,
    updateSessionUC *UpdateSessionUseCase) *LoginUseCase{

	return &LoginUseCase{
		jwtService:jwt,
		repo: repo,
		authService: authService,
        registerUseCase: registerUC,
        findUserByEmailUC:findUserByEmailUC,
        updateSessionUC:updateSessionUC,
	}
}

func ( l* LoginUseCase) Execute(token string, userAgent string) (*authentities.SessionEntity,error){
	// 1. Verificación externa
    userProvider, err := l.authService.VerifyAndExtract(token)
    if err != nil {
        return nil, err
    }

    refreshToken, _ := l.jwtService.GenerateToken(userProvider.Email, "72h")

    // 2. Intento de búsqueda
    user, err := l.findUserByEmailUC.Execute(userProvider.Email)
    
    // 3. Manejo de Usuario No Existente (Registro)
    // Si el error es 404, registramos y cortamos esa rama devolviendo la sesión
    if exceptions.IsNotFound(err) {
        return l.handleNewUser(userProvider, refreshToken, userAgent)
    }

    // 4. Si hubo otro error (500, etc.), salimos
    if err != nil {
        return nil, err
    }

    // 5. Usuario Existente: Actualizamos sesión
    if err := l.updateSessionUC.Execute(refreshToken, user.Id); err != nil {
        return nil, err
    }

    return l.buildSessionResponse(user, refreshToken)

}
func (l *LoginUseCase) buildSessionResponse(user *authentities.UserEntity, refreshToken string) (*authentities.SessionEntity, error) {
	accessToken, err := l.jwtService.GenerateToken(user.Email, "1m")
	if err != nil {
		return nil, err // O un AppError de dominio si prefieres
	}

	return &authentities.SessionEntity{
		UserData:     *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (l *LoginUseCase) handleNewUser(providerUser *authports.ResponseProvider, refreshToken string, userAgent string) (*authentities.SessionEntity, error) {
	newUserVO := authvalueobjects.CreateUserVO{
		ExternalID:   providerUser.ClerkId,
		Email:        providerUser.Email,
		PhotoUrl:     providerUser.PhotoUrl,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
	}

	userCreated, err := l.registerUseCase.Execute(newUserVO)
	if err != nil {
		return nil, err
	}

	return l.buildSessionResponse(userCreated, refreshToken)
}