package sharedHttp

import (
	"encoding/json"
	"net/http"

	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, err error) {
	// Intentamos castear al tipo AppError que creamos
	if appErr, ok := err.(*domainerrors.AppError); ok {
		RespondWithJSON(w, appErr.Status, map[string]any{
			"title":   appErr.Title,
			"message": appErr.Message,
			"status":  appErr.Status,
		})

	}

	// Error genérico para errores que no controlamos (500)
	RespondWithJSON(w, http.StatusInternalServerError, map[string]string{
		"title":   "Internal Server Error",
		"message": "An unexpected error occurred",
	})
}