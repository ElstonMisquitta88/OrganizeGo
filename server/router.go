package server

import (
	"OrganizeGo/handlers"
	"net/http"
)

// NewRouter configures all API routes and returns an http.Handler.
func NewRouter(h *handlers.TodoHandler) http.Handler {
	mux := http.NewServeMux()

	// /todos -> list + create
	mux.HandleFunc("/todos", h.HandleTodos)

	// /todos/{id} -> get, update, delete
	mux.HandleFunc("/todos/", h.HandleTodoByID)

	return mux
}
