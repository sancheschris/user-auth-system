package auth

import (
	"encoding/json"
	"net/http"

	jwtmiddleware "github.com/sancheschris/user-auth-system/internal/middleware"
	"github.com/sancheschris/user-auth-system/internal/model"
)

type AuthHandler struct {
	AuthService *AuthService
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}	

	if err := h.AuthService.Register(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	return
	}

	token, err := h.AuthService.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *AuthHandler) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
    userID, ok := jwtmiddleware.GetUserIDFromContext(r.Context())
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Welcome to the protected endpoint!",
        "user_id": userID,
    })
}