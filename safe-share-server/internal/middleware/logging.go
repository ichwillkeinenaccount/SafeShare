package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// wrappedWriter is a custom HTTP response writer that captures the status code of the response.
//
// Fields:
//
//	ResponseWriter (http.ResponseWriter): The original HTTP response writer.
//	statusCode (int): The status code of the response.
type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and writes the header to the response.
//
// Parameters:
//
//	statusCode (int): The status code to be written to the response.
func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// Logging logs information about the request and response.
//
// Parameters:
//
//	next (http.Handler): The next HTTP handler to be called in the middleware chain.
//
// Returns:
//
//	http.Handler: A new HTTP handler that wraps the next handler with logging functionality.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		slog.Info("Request handled:",
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"user_agent", r.UserAgent(),
			"status_code", wrapped.statusCode,
			"duration", time.Since(start),
			"client_ip", r.RemoteAddr,
		)
	})
}
