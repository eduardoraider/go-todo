package handlers

import (
	"encoding/json"
	"github.com/eduardoraider/go-todo/entity"
	"net/http"
)

// GetTodosHandler returns a handler function for retrieving all todo items
func GetTodosHandler(todoList *entity.TodoList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := todoList.GetTodos()
		respondJSON(w, http.StatusOK, todos)
	}
}

// AddTodoHandler returns a handler function for adding a new todo item
func AddTodoHandler(todoList *entity.TodoList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newItem entity.TodoItem
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todoList.AddTodoItem(newItem)
		w.WriteHeader(http.StatusCreated)
	}
}

// UpdateTodoHandler returns a handler function for updating a todo item
func UpdateTodoHandler(todoList *entity.TodoList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatedItem entity.TodoItem
		err := json.NewDecoder(r.Body).Decode(&updatedItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todoList.UpdateTodoItem(updatedItem)
		w.WriteHeader(http.StatusOK)
	}
}

// DeleteTodoHandler returns a handler function for deleting a todo item
func DeleteTodoHandler(todoList *entity.TodoList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todoID struct {
			ID string `json:"id"`
		}
		err := json.NewDecoder(r.Body).Decode(&todoID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todoList.DeleteTodoItem(todoID.ID)
		w.WriteHeader(http.StatusOK)
	}
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
