package authadapters

import (
	"fmt"
	"context"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/purplesvage/moneka-ai/cmd/config"
	authports "github.com/purplesvage/moneka-ai/internal/auth/domain/ports"
)



type AuthProviderAdapter struct{}

func NewAuthProviderAdapter() authports.AuthProviderPort{
	clerkKey := config.Envs.ClerkKey
	clerk.SetKey(clerkKey)
	return  &AuthProviderAdapter{} 
}

func getCustomClaims(m map[string]any, key string) string{
	val, ok := m[key].(string)
    if !ok {
        return ""
    }
    return val
}

func (a *AuthProviderAdapter) VerifyAndExtract(token string)(*authports.ResponseProvider, error){
	claims, err := jwt.Verify(context.Background(), &jwt.VerifyParams{
		Token: token,
	})
	if err != nil {
		return nil, fmt.Errorf("token de sesión inválido: %w", err)
	}
	clerkID := claims.Subject
	customMap, ok := claims.Custom.(map[string]any)
	if !ok {
		customMap = make(map[string]any)
	}

	email := getCustomClaims(customMap,"email")
	photo := getCustomClaims(customMap, "picture")
    phone := getCustomClaims(customMap, "phone_number")

	response := &authports.ResponseProvider{
        ClerkId: clerkID,
        Email:   email,
    }

    if photo != "" {
        response.PhotoUrl = &photo
    }
    if phone != "" {
        response.PhoneNumber = &phone
    }

    return response, nil
}

