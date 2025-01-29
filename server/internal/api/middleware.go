package api

import (
	"github.com/spf13/viper"
	"log/slog"
	"net/http"
	"time"
)

// chainMiddlewares applies a series of middlewares to an HTTP handler.
// The middlewares are applied in reverse order, with the first middleware
// being the innermost and the last middleware being the outermost.
//
// Parameters:
//
//	handler (http.Handler): The original HTTP handler to which the middlewares will be applied.
//
// Returns:
//
//	http.Handler: The HTTP handler wrapped with the specified middlewares.
func chainMiddlewares(handler http.Handler) http.Handler {
	//handler = headerCheckMiddleware(handler)
	handler = setAccessControlAllowOriginMiddleware(handler)
	handler = loggingMiddleware(handler)

	return handler
}

// loggingMiddleware logs each request
//
// Parameters:
//
//	next (http.Handler): The next HTTP handler to be called in the middleware chain.
//
// Returns:
//
//	http.Handler: A new HTTP handler that wraps the next handler with logging functionality.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		slog.Info("Request received",
			"status", wrapped.status,
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
			"client_ip", r.RemoteAddr,
		)
	})
}

// headerCheckMiddleware checks for a specific header
//
// Parameters:
//
//	next (http.Handler): The next HTTP handler to be called in the middleware chain.
//
// Returns:
//
//	http.Handler: A new HTTP handler that wraps the next handler with logging functionality.
func headerCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom-Header") == "" {
			http.Error(w, "Missing X-Custom-Header", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// setAccessControlAllowOriginMiddleware sets the Access-Control-Allow-Origin header, if the server is in development mode
//
// Parameters:
//
//	next (http.Handler): The next HTTP handler to be called in the middleware chain.
//
// Returns:
//
//	http.Handler: A new HTTP handler that wraps the next handler with logging functionality.
func setAccessControlAllowOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if viper.Get("development") == true {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			next.ServeHTTP(w, r)
		}
	})
}
