package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/torexanovich/savelife/backend/internal/database"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
    authService *Service
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService *Service) *AuthHandler {
    return &AuthHandler{authService}
}

// CreateRequestHandler handles the creation of ambulance requests
func (h *Handler) CreateRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req database.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	// Get user ID from request context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "missing userID in request context", http.StatusInternalServerError)
		return
	}

	// Fetch user details from the database
	user, err := h.userService.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "failed to fetch user details", http.StatusInternalServerError)
		return
	}

	// Set user details in the ambulance request
	req.FullName = user.FullName
	req.DateOfBirth = user.DateOfBirth
	req.Illness = user.Illness
	req.Location = user.Location
	req.PhoneNumber = user.PhoneNumber

	// Create the ambulance request
	if err := h.requestService.CreateRequest(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

// Login handles user login request
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    token, err := ah.authService.AuthenticateUser(context.Background(), credentials.Username, credentials.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
