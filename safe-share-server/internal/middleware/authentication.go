package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

const AuthUserID = "middleware.auth.userID"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement authentication!
		randomUserId := uuid.New()
		ctx := context.WithValue(r.Context(), AuthUserID, randomUserId)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
