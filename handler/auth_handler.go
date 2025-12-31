package handler

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"mini-jira/auth"
	"mini-jira/service"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByEmail(req.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, _ := auth.GenerateToken(user.ID)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
