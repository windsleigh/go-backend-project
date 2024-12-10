package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-backend-project/internal/handlers"

	"github.com/joho/godotenv"
)

func initializeRoutes() {
	http.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HelloHandler)))
	http.Handle("/health", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HealthHandler)))
}

func startServer() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// Get the port from the environment or use a default value
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	fmt.Printf("Server is running at http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func main() {
	initializeRoutes()
	startServer()
}
