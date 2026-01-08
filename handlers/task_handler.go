package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"test_mini_jira/models"
	"test_mini_jira/repository"
	"test_mini_jira/utils"
)

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := strings.Trim(strings.TrimPrefix(r.URL.Path, "/tasks"), "/")

	switch r.Method {

	case http.MethodPost:
		var t models.Task
		json.NewDecoder(r.Body).Decode(&t)
		if err := h.Repo.Create(t); err != nil {
			utils.JSONError(w, 500, "db error")
			return
		}
		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:
		tasks, err := h.Repo.GetAll()
		if err != nil {
			utils.JSONError(w, 500, "db error")
			return
		}
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPut:
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.JSONError(w, 400, "invalid id")
			return
		}
		var t models.Task
		json.NewDecoder(r.Body).Decode(&t)
		h.Repo.Update(id, t)

	case http.MethodPatch:
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.JSONError(w, 400, "invalid id")
			return
		}
		var b struct {
			Status string `json:"status"`
		}
		json.NewDecoder(r.Body).Decode(&b)
		h.Repo.PatchStatus(id, b.Status)

	case http.MethodDelete:
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.JSONError(w, 400, "invalid id")
			return
		}
		h.Repo.Delete(id)

	default:
		utils.JSONError(w, 405, "method not allowed")
	}
}
