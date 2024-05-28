package auth

import (
	"net/http"
	"strings"
)

// Middleware wraps an http.Handler and authenticates requests
type Middleware struct {
	authService *Service
}

// NewMiddleware creates a new authentication middleware
func NewMiddleware(authService *Service) *Middleware {
	return &Middleware{authService}
}

// Authenticate is the middleware function to authenticate requests
func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		// Extract the token from "Bearer <token>"
		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		token = parts[1]

		// Verify the token
		_, err := m.authService.ValidateToken(token)
		if err != nil {
			http.Error(w, "unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
