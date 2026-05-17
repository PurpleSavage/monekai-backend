package http

import (
	"encoding/json"
	"net/http"
	"time"
	authdtos "github.com/purplesvage/moneka-ai/internal/auth/in/dtos"
	authusecases "github.com/purplesvage/moneka-ai/internal/auth/usecases"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
	sharedHttp "github.com/purplesvage/moneka-ai/internal/shared/in/http"
	"github.com/purplesvage/moneka-ai/internal/shared/privatemiddlewares"
	"github.com/purplesvage/moneka-ai/pkg/middlewares"
)


type AuthHandler struct {
    loginUseCase *authusecases.LoginUseCase
    authMiddleware *privatemiddlewares.AuthMiddleware
    findUserUseCase *authusecases.FindUserByEmailUseCase
    refreshTokenUseCase *authusecases.RefreshTokenUseCase
}

func NewAuthHandler(
    lu *authusecases.LoginUseCase,
    am *privatemiddlewares.AuthMiddleware,
    fu *authusecases.FindUserByEmailUseCase,
    ru *authusecases.RefreshTokenUseCase,
) *AuthHandler {
	return &AuthHandler{
        loginUseCase: lu,
        authMiddleware: am,
        findUserUseCase: fu,
        refreshTokenUseCase:ru,
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
        sharedHttp.RespondWithError(w, err)
        return
    }
    refreshToken:= session.RefreshToken
    responseSession:= authdtos.ResponseSessionDto{
        UserData:authdtos.UserResponseDto{
            ID:        session.UserData.Id,
            Email:     session.UserData.Email,
            PhotoURL:  session.UserData.PhotoUrl,
            CreatedAt: session.UserData.CreatedAt.Format(time.RFC3339),
            Credits:   session.UserData.Credits,
        },
        AccessToken: session.AccessToken,
    }
    sharedHttp.HandleCookie(w,refreshToken)
    sharedHttp.RespondWithJSON(w, http.StatusOK, responseSession)
}


func (h *AuthHandler)RefreshToken(w http.ResponseWriter, r *http.Request){
    email, ok:= r.Context().Value(privatemiddlewares.EmailContextKey).(string) 
    if !ok {
        sharedHttp.RespondWithError(
            w,
            domainerrors.NewAppError(
                401,
                "Unauthorized",
                "Email not found in context",
                nil,
            ),
        )
        return
    } 
    newtoken,err:= h.refreshTokenUseCase.Excute(email)
    if err!= nil{
        sharedHttp.RespondWithError(
            w,
            err,
        )
    }
    response := authdtos.RefreshTokenResponseDto{
        AccessToken: newtoken.Value(),
    }
    sharedHttp.RespondWithJSON(w, http.StatusOK, response)
}




func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request){
    email, ok := r.Context().Value(privatemiddlewares.EmailContextKey).(string)
    if !ok {
        sharedHttp.RespondWithError(
            w,
            domainerrors.NewAppError(
                401,
                "Unauthorized",
                "Email not found in context",
                nil,
            ),
        )
        return
    }
    session,err:=h.findUserUseCase.Execute(email)
    if err != nil {
        sharedHttp.RespondWithError(
            w,
            err,
        )
        return
    }
    user:=authdtos.UserResponseDto{
        ID:        session.Id,
        Email:     session.Email,
        PhotoURL:  session.PhotoUrl,
        CreatedAt: session.CreatedAt.Format(time.RFC3339),
        Credits:   session.Credits,
    }
    sharedHttp.RespondWithJSON(w, http.StatusOK, user)
}


func MapRoutes(mux *http.ServeMux, h *AuthHandler) {
    mdls := []middlewares.Middleware{
		h.authMiddleware.RefreshToken,
	}

    //ruta -> entrar o crear una cuenta 
	mux.HandleFunc("POST /login", h.Login)


    // ruta -> obtener perfil
	protectedHandler := middlewares.ContextMiddleware(
		http.HandlerFunc(h.GetProfile),
		mdls,
	)
	mux.Handle("GET /profile", protectedHandler)



    // ruta -> refrescar token
    protectedHandlerRefreshToken:=middlewares.ContextMiddleware(
		http.HandlerFunc(h.RefreshToken),
		mdls,
	)
	mux.Handle("GET /profile", protectedHandlerRefreshToken)
}