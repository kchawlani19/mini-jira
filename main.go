package main

import (
	"fmt"
	"mini-jira/handler"
	"mini-jira/repository"
	"mini-jira/service"
	"net/http"
)

func main() {

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/users", userHandler.CreateUser)

	fmt.Println("MINI_JIRA server running on :8080")
	http.ListenAndServe(":8080", nil)
}
