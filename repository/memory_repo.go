package repository

import (
	"OrganizeGo/models"
	"errors"
	"time"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type TodoRepository interface {
	List() ([]models.Todo, error)
	Get(id int) (models.Todo, error)
	Create(title string) (models.Todo, error)
}

type MemoryTodoRepo struct {
	todos  []models.Todo
	nextID int
}

// Explain what this function does
func NewMemoryTodoRepo() *MemoryTodoRepo {
	return &MemoryTodoRepo{
		todos:  []models.Todo{}, // empty slice to start
		nextID: 1,
	}
}

// (a) List of all todos
func (r *MemoryTodoRepo) List() ([]models.Todo, error) {
	return r.todos, nil
}

// (b) Get a todos by ID
func (r *MemoryTodoRepo) Get(id int) (models.Todo, error) {
	for _, t := range r.todos {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Todo{}, ErrTodoNotFound
}

// (c) Create a new todos
func (r *MemoryTodoRepo) Create(title string) (models.Todo, error) {
	t := models.Todo{
		ID:        r.nextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	r.nextID++
	r.todos = append(r.todos, t)

	return t, nil
}
