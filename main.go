package main

import (
	"log"
	"net/http"

	"OrganizeGo/handlers"
	"OrganizeGo/repository"
	"OrganizeGo/server"
)

func main() {
	// 1) Create repository
	repo := repository.NewMemoryTodoRepo()

	// 2) Create handler
	handler := handlers.NewTodoHandler(repo)

	// 3) Create router
	mux := server.NewRouter(handler)

	// 4) Start server
	log.Println("🚀 Todo API server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
