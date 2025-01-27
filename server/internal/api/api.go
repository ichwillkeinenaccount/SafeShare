package api

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func StartApi() {
	textHandler := &TextHandler{}

	router := initRoutes(textHandler)
	routerWithMiddlewares := chainMiddlewares(router)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: routerWithMiddlewares,
	}

	slog.Info(fmt.Sprintf("Server running on port %d", viper.GetInt("server.port")))
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Could not start api", "error", err)
		return
	}
}

func initRoutes(textHandler *TextHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/text/", textHandler.getAll)
	mux.HandleFunc("POST /api/v1/text/create", textHandler.postText)
	return mux
}

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
	// Apply headerCheckMiddleware first
	handler = headerCheckMiddleware(handler)
	// Apply loggingMiddleware next
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
			"headers", r.Header,
			"client_ip", r.RemoteAddr,
		)
	})
}

// headerCheckMiddleware checks for a specific header
func headerCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom-Header") == "" {
			http.Error(w, "Missing X-Custom-Header", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
