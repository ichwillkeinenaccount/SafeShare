package middleware

import (
	"github.com/spf13/viper"
	"net/http"
)

// AllowCors sets the Access-Control-Allow headers, if the server is in development mode.
//
// Parameters:
//
//	next (http.Handler): The next HTTP handler to be called in the middleware chain.
//
// Returns:
//
//	http.Handler: A new HTTP handler that wraps the next handler with logging functionality.
func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if viper.Get("development") == true {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		}
	})
}
