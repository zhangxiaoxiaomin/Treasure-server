package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"treasure-server/database"
	"treasure-server/handlers"
	"treasure-server/repository"
)

func main() {
	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize database
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/treasure.db"
	}

	if err := database.Init(dbPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database initialized successfully")

	// Initialize repository and handler
	repo := repository.NewCollectionRepo()
	handler := handlers.NewCollectionHandler(repo)

	// Register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/collections", handler.HandleCollections)
	mux.HandleFunc("/api/collections/", handler.HandleCollections)

	// Health check endpoint
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"treasure-server"}`))
	})

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Treasure server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}