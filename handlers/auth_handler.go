package handlers

import (
	"encoding/json"
	"net/http"

	"test_mini_jira/models"
	"test_mini_jira/repository"
	"test_mini_jira/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Repo *repository.UserRepository
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.JSONError(w, 400, "invalid json")
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hash)
	if u.Role == "" {
		u.Role = "USER"
	}

	if err := h.Repo.Create(u); err != nil {
		utils.JSONError(w, 500, "email already exists")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.Repo.FindByEmail(req.Email)
	if err != nil {
		utils.JSONError(w, 401, "invalid credentials")
		return
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(req.Password),
	) != nil {
		utils.JSONError(w, 401, "invalid credentials")
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
