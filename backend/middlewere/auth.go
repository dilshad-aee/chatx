package middlewere

import (
	"chatx/backend/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	userIDKey   contextKey = "userID"
	usernameKey contextKey = "username"
)

func AuthMiddlewere(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {

			http.Error(w, "Missing or invaild Authorization header", http.StatusUnauthorized)
			return

		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenString)

		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		ctx = context.WithValue(ctx, usernameKey, claims.Username)

		next(w, r.WithContext(ctx))

	}
}
