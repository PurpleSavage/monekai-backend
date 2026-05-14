package middlewares

import "net/http"

func ContextMiddleware(next http.Handler, middlewares []func(next http.Handler) http.Handler) http.Handler {
	finalHandler := next

	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i](finalHandler)
	}

	return finalHandler
}