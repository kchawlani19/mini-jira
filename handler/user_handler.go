package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"mini-jira/service"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		h.handleCreate(w, r)

	case http.MethodGet:
		if r.URL.Path == "/users" {
			h.handleGetAll(w, r)
		} else {
			h.handleGetByID(w, r)
		}

	case http.MethodPut:
		h.handleUpdate(w, r)

	case http.MethodDelete:
		h.handleDelete(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) handleCreate(w http.ResponseWriter, r *http.Request) {

	var req CreateUserRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.service.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) handleGetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetAllUsers())
}

func (h *UserHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[2])

	user, err := h.service.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[2])

	var req CreateUserRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.service.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) handleDelete(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[2])

	if err := h.service.DeleteUser(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
