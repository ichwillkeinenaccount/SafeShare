package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// CreateStack applies a series of middlewares to an HTTP handler.
// The middlewares are applied in order, with the first middleware
// being the outermost and the last middleware being the innermost.
//
// Parameters:
//
//	handler (http.Handler): The original HTTP handler to which the middlewares will be applied.
//
// Returns:
//
//	http.Handler: The HTTP handler wrapped with the specified middlewares.
func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}
