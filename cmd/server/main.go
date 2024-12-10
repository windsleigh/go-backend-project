package main

import (
	"fmt"
	"go-backend-project/internal/handlers" // Import the handlers package
	"net/http"
)

// initializeRoutes sets up the routes for the server.
func initializeRoutes() {
	http.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HelloHandler)))
	http.Handle("/health", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HealthHandler)))
}

// startServer starts the HTTP server.
func startServer() {
	port := ":8080"
	fmt.Printf("Server is running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func main() {
	initializeRoutes() // Set up routes
	startServer()      // Start the server
}
