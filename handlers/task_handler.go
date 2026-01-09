package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"test_mini_jira/middleware"
	"test_mini_jira/models"
	"test_mini_jira/repository"
	"test_mini_jira/utils"
)

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := strings.Trim(strings.TrimPrefix(r.URL.Path, "/tasks"), "/")

	userID := r.Context().Value(middleware.UserIDKey).(int)
	role := r.Context().Value(middleware.RoleKey).(string)

	switch r.Method {

	case http.MethodPost:
		var t models.Task
		json.NewDecoder(r.Body).Decode(&t)

		// USER cannot assign task to someone else
		if role != "ADMIN" && t.AssigneeID != nil && *t.AssigneeID != userID {
			utils.JSONError(w, 403, "only admin can assign tasks")
			return
		}

		h.Repo.Create(t)
		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:
		tasks, _ := h.Repo.GetAll()
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPut:
		id, _ := strconv.Atoi(idStr)

		if role != "ADMIN" {
			ok, _ := h.Repo.IsOwner(id, userID)
			if !ok {
				utils.JSONError(w, 403, "not allowed to update this task")
				return
			}
		}

		var t models.Task
		json.NewDecoder(r.Body).Decode(&t)
		h.Repo.Update(id, t)

	case http.MethodPatch:
		id, _ := strconv.Atoi(idStr)

		if role != "ADMIN" {
			ok, _ := h.Repo.IsOwner(id, userID)
			if !ok {
				utils.JSONError(w, 403, "not allowed to update this task")
				return
			}
		}

		var b struct {
			Status string `json:"status"`
		}
		json.NewDecoder(r.Body).Decode(&b)
		h.Repo.PatchStatus(id, b.Status)

	case http.MethodDelete:
		id, _ := strconv.Atoi(idStr)

		if role != "ADMIN" {
			ok, _ := h.Repo.IsOwner(id, userID)
			if !ok {
				utils.JSONError(w, 403, "not allowed to delete this task")
				return
			}
		}

		h.Repo.Delete(id)

	default:
		utils.JSONError(w, 405, "method not allowed")
	}
}
