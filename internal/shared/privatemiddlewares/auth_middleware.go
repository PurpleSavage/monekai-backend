package privatemiddlewares

import (
	"net/http"
	"strings"

	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
	sharedHttp "github.com/purplesvage/moneka-ai/internal/shared/in/http"
)

func AuthMiddleware(next http.Handler) http.Handler{
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            sharedHttp.RespondWithError(w, domainerrors.NewAppError(401, "Unauthorized", "Missing or invalid token", nil))
            return
        }

        // 2. Si todo está bien, llamar al siguiente handler
        next.ServeHTTP(w, r)
	})
}

