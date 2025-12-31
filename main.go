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

	// PUBLIC ROUTES
	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/users", userHandler.CreateUser) // POST allowed without token

	// PROTECTED ROUTES
	http.HandleFunc("/users/", middleware.JWTAuth(userHandler.CreateUser))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
