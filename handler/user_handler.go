package handler

import (
	"encoding/json"
	"mini-jira/service"
	"net/http"
	"strconv"
	"strings"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json"email"`
}

type CreateUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.handleCreate(w, r)
		return
	}

	if r.Method == http.MethodGet {
		if r.URL.Path == "/users" {
			h.handleGetAll(w, r)
			return
		}

		h.handleGetByID(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *UserHandler) handleCreate(w http.ResponseWriter, r *http.Request) {

	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(req.Name, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp := CreateUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) handleGetAll(w http.ResponseWriter, r *http.Request) {

	users := h.service.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
