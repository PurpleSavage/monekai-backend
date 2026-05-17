package audioinadapters
import (
	"sync"
)


type SSEManager struct {
	// Mapeamos el Email del usuario con su canal de strings (o structs)
	mu  sync.RWMutex
	clients map[string]chan string
}

func NewSSEManager() *SSEManager {
	return &SSEManager{
		clients: make(map[string]chan string),
	}
}

// Register crea un canal para el usuario cuando se conecta al SSE
func (m *SSEManager) Register(email string) chan string {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Si ya existía una conexión vieja, la cerramos
	if oldChan, exists := m.clients[email]; exists {
		close(oldChan)
	}

	ch := make(chan string, 10) // Buffer de 10 mensajes
	m.clients[email] = ch
	return ch
}

// Unregister limpia el canal cuando el usuario cierra la pestaña
func (m *SSEManager) Unregister(email string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if ch, exists := m.clients[email]; exists {
		close(ch)
		delete(m.clients, email)
	}
}

// BroadcastToUser envía datos de forma segura al usuario específico desde el Webhook
func (m *SSEManager) BroadcastToUser(email string, data string) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if ch, exists := m.clients[email]; exists {
		ch <- data // Enviamos el JSON al canal del usuario
	}
}