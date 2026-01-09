package handlers

import (
	"database/sql"
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

	// ✅ CREATE TASK (UNASSIGNED BY DEFAULT)
	case http.MethodPost:
		var t models.Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			utils.JSONError(w, 400, "invalid JSON")
			return
		}

		if strings.TrimSpace(t.Title) == "" {
			utils.JSONError(w, 400, "title is required")
			return
		}

		// 🔐 USER cannot assign
		if role != "ADMIN" && t.AssigneeID != nil {
			utils.JSONError(w, 403, "only admin can assign tasks")
			return
		}

		h.Repo.Create(t)
		w.WriteHeader(http.StatusCreated)

	// ✅ GET TASKS (ROLE + PAGINATION)
	case http.MethodGet:
		q := r.URL.Query()

		page, _ := strconv.Atoi(q.Get("page"))
		limit, _ := strconv.Atoi(q.Get("limit"))

		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		offset := (page - 1) * limit

		var tasks []models.Task
		var err error

		if role == "ADMIN" {
			tasks, err = h.Repo.GetAll(limit, offset)
		} else {
			tasks, err = h.Repo.GetByUser(userID, limit, offset)
		}

		if err != nil {
			utils.JSONError(w, 500, "failed to fetch tasks")
			return
		}

		json.NewEncoder(w).Encode(tasks)

	// PUT / PATCH / DELETE (OWNERSHIP + ADMIN OVERRIDE)
	case http.MethodPut, http.MethodPatch, http.MethodDelete:
		id, _ := strconv.Atoi(idStr)

		if role != "ADMIN" {
			ok, _ := h.Repo.IsOwner(id, userID)
			if !ok {
				utils.JSONError(w, 403, "not allowed to modify this task")
				return
			}
		}

		if r.Method == http.MethodDelete {
			err := h.Repo.Delete(id)
			if err == sql.ErrNoRows {
				utils.JSONError(w, 404, "task not found")
				return
			}

			return
		}

		if r.Method == http.MethodPatch {
			var b struct {
				Status string `json:"status"`
			}
			json.NewDecoder(r.Body).Decode(&b)
			h.Repo.PatchStatus(id, b.Status)
			return
		}

		var t models.Task
		json.NewDecoder(r.Body).Decode(&t)
		h.Repo.Update(id, t)

	default:
		utils.JSONError(w, 405, "method not allowed")
	}
}
