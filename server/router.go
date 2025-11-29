package server

import (
	"OrganizeGo/handlers"
	"net/http"
)

// NewRouter configures all API routes and returns an http.Handler.
func NewRouter(h *handlers.TodoHandler) http.Handler {
	mux := http.NewServeMux()

	// /todos/ListTodos -> List
	mux.HandleFunc("/todos/ListTodos", h.ListTodos)

	// /todos/Create -> Create
	mux.HandleFunc("/todos/Create", h.CreateTodo)

	// /todos/{id} -> Fetch by Id
	mux.HandleFunc("/todos/", h.HandleTodoByID)

	return mux
}
