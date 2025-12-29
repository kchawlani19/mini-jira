package main

import (
	"fmt"
	"mini-jira/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handler.LoginHandler)

	fmt.Println("MINI_JIRA server running on :8080")
	http.ListenAndServe(":8080", nil)
}
