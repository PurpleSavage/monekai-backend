package http

import (
	"encoding/json"
	"net/http"

	audioinadapters "github.com/purplesvage/moneka-ai/internal/audio/in/adapters"
	audioincommands "github.com/purplesvage/moneka-ai/internal/audio/in/dtos/commands"
	audiousecases "github.com/purplesvage/moneka-ai/internal/audio/usecases"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
	sharedHttp "github.com/purplesvage/moneka-ai/internal/shared/in/http"
	sharedvalidators "github.com/purplesvage/moneka-ai/internal/shared/in/validators"
	"github.com/purplesvage/moneka-ai/internal/shared/privatemiddlewares"
	"github.com/purplesvage/moneka-ai/pkg/middlewares"
)

type AudioHandler struct {
	authMiddleware *privatemiddlewares.AuthMiddleware
	generateSongUseCase *audiousecases.GenerateSongUseCase
	validator           *sharedvalidators.DTOValidator
	sseManager *audioinadapters.SSEManager
}

func NewAudioHandler(
	am *privatemiddlewares.AuthMiddleware,
	gs *audiousecases.GenerateSongUseCase, 
	v *sharedvalidators.DTOValidator,
	sse *audioinadapters.SSEManager,
) *AudioHandler {
	return &AudioHandler{
		authMiddleware: am,
		generateSongUseCase:gs,
		validator:           v,
		sseManager:sse,
	}
}


func (h *AudioHandler) CreateSong(w http.ResponseWriter, r *http.Request){
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

	var dto audioincommands.GenerateSongDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		sharedHttp.RespondWithError(
			w,
			domainerrors.NewAppError(
				http.StatusBadRequest,
				"Bad Request",
				"Malformed JSON payload",
				err,
			),
		)
		return
	}
	if err := h.validator.ValidateStruct(dto); err != nil {
		sharedHttp.RespondWithError(w, err)
		return
	}
	response, err := h.generateSongUseCase.Execute(dto, email)
	if err != nil {
		sharedHttp.RespondWithError(w, err)
		return
	}
	sharedHttp.RespondWithJSON(w, http.StatusOK, response)

}




func MapRoutes(mux *http.ServeMux, h *AudioHandler) {
	mdls := []middlewares.Middleware{
		h.authMiddleware.RefreshToken,
	}
	protectedHandler := middlewares.ContextMiddleware(
		http.HandlerFunc(h.CreateSong),
		mdls,
	)
	mux.Handle("POST /create", protectedHandler)

}