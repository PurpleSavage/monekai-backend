package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler
func ContextMiddleware(next http.Handler, middlewares []Middleware ) http.Handler {
	finalHandler := next

	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i](finalHandler)
	}

	return finalHandler
}