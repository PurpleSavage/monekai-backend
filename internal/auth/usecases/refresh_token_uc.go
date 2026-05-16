package authusecases

import (
	authvalueobjects "github.com/purplesvage/moneka-ai/internal/auth/domain/valueobjects"
	sharedports "github.com/purplesvage/moneka-ai/internal/shared/domain/ports"
)


type RefreshTokenUseCase struct{
	jwtService  sharedports.JwtPort
}

func NewRefreshTokenUseCase(
	jwtService sharedports.JwtPort, 
) *RefreshTokenUseCase{
	return &RefreshTokenUseCase{
		jwtService: jwtService,
	}
}


func (r *RefreshTokenUseCase) Excute(email string)(authvalueobjects.TokenVO, error){
	token, err:= r.jwtService.GenerateToken(email, "1m")
	if err!= nil{
		return  authvalueobjects.TokenVO{},err
	}
	validatedtoken,err:= authvalueobjects.NewTokenVO(token)
	if err!= nil{
		return  authvalueobjects.TokenVO{},err
	}
	return validatedtoken,nil
}
