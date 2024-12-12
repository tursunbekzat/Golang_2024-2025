package middlewares

import (
	"backend/auth"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey = contextKey("user")

// AuthMiddleware validates the JWT token and extracts user information
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, claims, err := auth.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add claims to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext retrieves user claims from the request context
func GetUserFromContext(r *http.Request) map[string]interface{} {
	claims, ok := r.Context().Value(UserContextKey).(map[string]interface{})
	if !ok {
		return nil
	}
	return claims
}

// RoleMiddleware restricts access based on user role
func RoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := GetUserFromContext(r)
			if claims == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			role, ok := claims["role"].(string)
			if !ok || role != requiredRole {
				http.Error(w, "Access forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
