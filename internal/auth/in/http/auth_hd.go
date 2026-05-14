package http

import (
	"encoding/json"
	"net/http"

	authdtos "github.com/purplesvage/moneka-ai/internal/auth/in/dtos"
	authusecases "github.com/purplesvage/moneka-ai/internal/auth/usecases"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
	sharedHttp "github.com/purplesvage/moneka-ai/internal/shared/in/http"
	"github.com/purplesvage/moneka-ai/internal/shared/privatemiddlewares"
	"github.com/purplesvage/moneka-ai/pkg/middlewares"
)


type AuthHandler struct {
    loginUseCase *authusecases.LoginUseCase
  
}

func NewAuthHandler(lu *authusecases.LoginUseCase) *AuthHandler {
	return &AuthHandler{
        loginUseCase: lu,
    }
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req authdtos.AuthRequestDto
    userAgent:=r.UserAgent()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        sharedHttp.RespondWithError(w, domainerrors.NewAppError(400, "Bad Request", "Invalid JSON body", err))
        return
    }

	if req.Token == "" {
        sharedHttp.RespondWithError(w, domainerrors.NewAppError(400, "Validation Error", "Token is required", nil))
        return
    }
    if userAgent==""{
        userAgent = "unknown"
    }
    session, err := h.loginUseCase.Execute(req.Token,userAgent) 
    if err != nil {
        // El helper se encarga de loguear y poner el status correcto (401, 404, 500)
        sharedHttp.RespondWithError(w, err)
        return
    }
    refreshToken:= session.RefreshToken
    responseSession:= &authdtos.ResponseSessionDto{
        UserData:session.UserData,
        AccessToken: session.AccessToken,
    }
    // manejar la cookie de session
    sharedHttp.HandleCookie(w,refreshToken)
    // 3. Retornar la sesión exitosa en formato JSON
    sharedHttp.RespondWithJSON(w, http.StatusOK, responseSession)
}


func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request){

}


func MapRoutes(mux *http.ServeMux, h *AuthHandler) {
	mux.HandleFunc("POST /login", h.Login)

    mdls := []func(http.Handler) http.Handler{
		privatemiddlewares.AuthMiddleware,
	}

	protectedHandler := middlewares.ContextMiddleware(
		http.HandlerFunc(h.GetProfile),
		mdls,
	)

	mux.Handle("POST /profile", protectedHandler)
}