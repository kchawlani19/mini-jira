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

	// public
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	// protected
	http.HandleFunc("/tasks", middleware.JWTAuth(taskHandler.ServeHTTP))
	http.HandleFunc("/tasks/", middleware.JWTAuth(taskHandler.ServeHTTP))

	http.ListenAndServe(":8080", nil)
}
