package privatemiddlewares

import (
	"context"
	"net/http"
	"strings"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
	sharedports "github.com/purplesvage/moneka-ai/internal/shared/domain/ports"
	sharedHttp "github.com/purplesvage/moneka-ai/internal/shared/in/http"
)

type Middleware func(http.Handler) http.Handler
type contextKey string

const EmailContextKey contextKey = "email"

type AuthMiddleware struct{
	JwtService sharedports.JwtPort
}
func NewAuthMiddleware(JwtService sharedports.JwtPort) *AuthMiddleware {
	return &AuthMiddleware{
		JwtService: JwtService,
	}
}
func (a *AuthMiddleware) AccesToken(next http.Handler)http.Handler{
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            sharedHttp.RespondWithError(w, domainerrors.NewAppError(401, "Unauthorized", "Missing or invalid token", nil))
            return
        }
		token:= strings.Split(authHeader," ")[1]
		email, err:= a.JwtService.VerifyToken(token)
		if err != nil {
			sharedHttp.RespondWithError(w, domainerrors.NewAppError(401, "Unauthorized", "token expired", nil))
			return 
		}
		ctx := context.WithValue(
			r.Context(),
			EmailContextKey,
			email,
		)
        // 2. Si todo está bien, llamar al siguiente handler
        next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func (a *AuthMiddleware) RefreshToken(next http.Handler)http.Handler{
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cookie,err:= r.Cookie("session_token")
		if err!= nil {
			sharedHttp.RespondWithError(w, domainerrors.NewAppError(401, "Unauthorized", "Missing or invalid refresh token", nil))
			return
		}
		email, err:= a.JwtService.VerifyToken(cookie.Value)
		if err != nil {
			sharedHttp.RespondWithError(
				w,
				domainerrors.NewAppError(
					401,
					"Unauthorized",
					"Expired refresh token",
					nil,
				),
			)
			return
		}
		ctx := context.WithValue(
			r.Context(),
			EmailContextKey,
			email,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}



