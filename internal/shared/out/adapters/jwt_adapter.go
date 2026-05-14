package sharedadapters

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/purplesvage/moneka-ai/cmd/config"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
	sharedports "github.com/purplesvage/moneka-ai/internal/shared/domain/ports"
)

type JwtAdapterService struct{}

func NewJwtAdapterService() sharedports.JwtPort{
	return  &JwtAdapterService{}
}

func (a *JwtAdapterService) GenerateToken(email string, durationStr string) (string, error) {
	secret := config.Envs.SecretJwt
	duration, err := time.ParseDuration(durationStr)
	
	if err != nil {
		// 400 Bad Request: The provided duration string is malformed
		return "", domainerrors.NewAppError(400, "Invalid Duration", "The token duration format is incorrect", err)
	}

	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(duration).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
    
	if err != nil {
		// 500 Internal Server Error: Cryptographic signing failed
		return "", domainerrors.NewAppError(500, "Signature Error", "Could not sign the security token", err)
	}

	return signedToken, nil
}

func (a *JwtAdapterService) VerifyToken(token string) (string, error) {
	secret := config.Envs.SecretJwt

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		// 401 Unauthorized: Token is expired, malformed, or signature is invalid
		return "", domainerrors.NewAppError(401, "Invalid Token", "Session has expired or token is corrupt", err)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", domainerrors.NewAppError(
			401,
			"Unauthorized",
			"Token content is not processable",
			nil,
		)
	}
	email, ok := claims["sub"].(string)
	if !ok {
		return "", domainerrors.NewAppError(
			401,
			"Unauthorized",
			"Invalid token payload",
			nil,
		)
	}

	return email, nil
}