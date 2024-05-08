package main

import (
	"fmt"
	"github.com/eduardoraider/go-todo/entity"
	"github.com/eduardoraider/go-todo/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialize todoList with some dummy data
	todoList := entity.NewTodoList()

	// Define routes
	http.HandleFunc("/todos", handlers.GetTodosHandler(todoList))
	http.HandleFunc("/todos/add", handlers.AddTodoHandler(todoList))
	http.HandleFunc("/todos/update", handlers.UpdateTodoHandler(todoList))
	http.HandleFunc("/todos/delete", handlers.DeleteTodoHandler(todoList))

	// Start server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
