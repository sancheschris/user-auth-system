package auth

import (
	"net/http"
)

type Handler struct {
	AuthService AuthService
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// lógica de login
}