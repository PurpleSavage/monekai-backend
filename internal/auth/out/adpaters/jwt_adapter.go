package authadapters

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/purplesvage/moneka-ai/cmd/config"
	authports "github.com/purplesvage/moneka-ai/internal/auth/domain/ports"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
)

type JwtAdapterService struct{}

func NewJwtAdapterService() authports.JwtPort{
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

func (a *JwtAdapterService) VerifyToken(token string) (jwt.MapClaims, error) {
	secret := config.Envs.SecretJwt

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		// 401 Unauthorized: Token is expired, malformed, or signature is invalid
		return nil, domainerrors.NewAppError(401, "Invalid Token", "Session has expired or token is corrupt", err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, domainerrors.NewAppError(401, "Unauthorized", "Token content is not processable", nil)
}