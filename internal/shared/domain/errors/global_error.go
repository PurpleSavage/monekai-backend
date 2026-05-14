package domainerrors

import "fmt"

type AppError struct {
    Title   string
    Message string
    Status  int // Ej: 404, 400, 500
    Err     error // El error original (para debugging)
}

// Implementamos la interfaz 'error' de Go
func (e *AppError) Error() string {
    return fmt.Sprintf("[%d] %s: %s", e.Status, e.Title, e.Message)
}

func NewAppError(status int, title, message string, err error) *AppError {
    return &AppError{
        Status:  status,
        Title:   title,
        Message: message,
        Err:     err,
    }
}