package main

import (
	"net/http"

	"test_mini_jira/db"
	"test_mini_jira/handlers"
	"test_mini_jira/middleware"
	"test_mini_jira/repository"
)

func main() {
	database := db.Connect()

	userRepo := &repository.UserRepository{DB: database}
	taskRepo := &repository.TaskRepository{DB: database}

	authHandler := &handlers.AuthHandler{Repo: userRepo}
	taskHandler := &handlers.TaskHandler{Repo: taskRepo}
	userHandler := &handlers.UserHandler{Repo: userRepo}

	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	// ✅ STRUCT-BASED HANDLERS
	http.Handle("/tasks", middleware.JWTAuth(taskHandler))
	http.Handle("/tasks/", middleware.JWTAuth(taskHandler))
	http.Handle("/users", middleware.JWTAuth(userHandler))

	http.ListenAndServe(":8080", nil)
}
