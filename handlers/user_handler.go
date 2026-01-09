package handlers

import (
	"encoding/json"
	"net/http"

	"test_mini_jira/middleware"
	"test_mini_jira/repository"
	"test_mini_jira/utils"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middleware.RoleKey).(string)

	if role != "ADMIN" {
		utils.JSONError(w, 403, "admin access only")
		return
	}

	if r.Method != http.MethodGet {
		utils.JSONError(w, 405, "method not allowed")
		return
	}

	users, err := h.Repo.GetAll()
	if err != nil {
		utils.JSONError(w, 500, "failed to fetch users")
		return
	}

	json.NewEncoder(w).Encode(users)
}
