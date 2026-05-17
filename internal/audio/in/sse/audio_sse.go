package audiosse


import (
	"fmt"
	"net/http"
	"time"
	"github.com/purplesvage/moneka-ai/pkg/middlewares"
	"github.com/purplesvage/moneka-ai/internal/audio/in/adapters" // Ajusta esta ruta a tu SSEManager
	"github.com/purplesvage/moneka-ai/internal/shared/privatemiddlewares"
)

// 1. Definimos la estructura propia para este Handler
type AudioSSEHandler struct {
	sseManager *audioinadapters.SSEManager
	authMiddleware *privatemiddlewares.AuthMiddleware
}

// 2. El constructor independiente
func NewAudioSSEHandler(
	sm *audioinadapters.SSEManager,
	am *privatemiddlewares.AuthMiddleware,
) *AudioSSEHandler {
	return &AudioSSEHandler{
		sseManager: sm,
	}
}

// 3. El método que maneja el streaming de eventos
func (h *AudioSSEHandler) StreamSongStatus(w http.ResponseWriter, r *http.Request) {
	// Extraemos el email del contexto de forma segura gracias a tu middleware de JWT
	email, ok := r.Context().Value(privatemiddlewares.EmailContextKey).(string)
	if !ok {
		http.Error(w, "Unauthorized: Email not found in context", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Registramos al usuario en nuestro gestor usando su email real
	userChan := h.sseManager.Register(email)
	defer h.sseManager.Unregister(email)

	keepAliveTicker := time.NewTicker(15 * time.Second)
	defer keepAliveTicker.Stop()

	clientGone := r.Context().Done()

	for {
		select {
		case <-clientGone:
			return 

		case <-keepAliveTicker.C:
			fmt.Fprintf(w, ": keep-alive\n\n")
			flusher.Flush()

		case eventData := <-userChan:
			fmt.Fprintf(w, "data: %s\n\n", eventData)
			flusher.Flush()
		}
	}
}

// 4. Su propia función para mapear las rutas de este Handler específico
func MapSSERoutes(mux *http.ServeMux, h *AudioSSEHandler) {


	mdls := []middlewares.Middleware{
		h.authMiddleware.RefreshToken,
	}
	protectedHandler := middlewares.ContextMiddleware(
		http.HandlerFunc(h.StreamSongStatus),
		mdls,
	)

	// Si usas tu función Chain que armamos antes, puedes encadenarlo de forma limpia:
	// protectedSSE := sharedHttp.Chain(ssePipeline, authMiddleware.AccessToken)
	
	mux.Handle("GET /audio/stream", protectedHandler)
}