package handlers

import (
	"OrganizeGo/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TodoHandler handles HTTP requests for todos.
// like a private field inside the class in C#
type TodoHandler struct {
	repo repository.TodoRepository
}

// Dependency Injection + Loose Coupling
// NewTodoHandler creates a new handler with the given repository.
func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

/* // HandleTodos routes based on method for /todos
func (h *TodoHandler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listTodos(w, r)
	case http.MethodPost:
		h.createTodo(w, r)
	default:
		methodNotAllowed(w, r)
	}
}



} */
// HandleTodoByID routes based on method for /todos/{id}
func (h *TodoHandler) HandleTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	switch r.Method {
	case http.MethodGet:
		h.getTodo(w, r, id)
	default:
		methodNotAllowed(w, r)
	}
}

// listTodos handles GET /todos
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.List()
	if err != nil {
		log.Printf("list error: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusOK, todos)
}

// createTodo handles POST /todos
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if strings.TrimSpace(req.Title) == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	todo, err := h.repo.Create(req.Title)
	if err != nil {
		log.Printf("create error: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Location", "/todos/"+strconv.Itoa(todo.ID))
	writeJSON(w, http.StatusCreated, todo)
}

// getTodo handles GET /todos/{id}
func (h *TodoHandler) getTodo(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := h.repo.Get(id)
	if err != nil {
		if err == repository.ErrTodoNotFound {
			writeError(w, http.StatusNotFound, "todo not found")
			return
		}
		log.Printf("get error: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

// ===== Helper functions =====

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON error: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	writeJSON(w, status, errorResponse{Error: message})
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}

// extractIDFromPath parses the ID from /todos/{id}
func extractIDFromPath(path string) (int, error) {
	// e.g. "/todos/123"
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) != 2 {
		return 0, strconv.ErrSyntax
	}
	return strconv.Atoi(parts[1])
}
