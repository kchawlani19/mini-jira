package main

import (
	"fmt"
	"net/http"

	"mini-jira/handler"
	"mini-jira/middleware"
	"mini-jira/repository"
	"mini-jira/service"
)

func main() {

	db := repository.NewMySQLDB()

	var userRepo repository.UserRepository
	userRepo = repository.NewMySQLUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService)

	taskRepo := repository.NewMySQLTaskRepository(db)
	taskService := service.NewTaskService(taskRepo, userRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// PUBLIC ROUTES
	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/users", userHandler.CreateUser) // POST allowed without token

	// PROTECTED ROUTES
	http.HandleFunc("/users/", middleware.JWTAuth(userHandler.CreateUser))

	http.HandleFunc("/tasks", taskHandler.CreateTask)       // POST
	http.HandleFunc("/tasks/", taskHandler.GetTaskByID)     // GET /tasks/{id}
	http.HandleFunc("/tasks-list", taskHandler.GetAllTasks) // GET all

	http.HandleFunc("/tasks/assign/", taskHandler.AssignTask)
	http.HandleFunc("/tasks/status/", taskHandler.UpdateTaskStatus)
	http.HandleFunc("/tasks/delete/", taskHandler.DeleteTask)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
