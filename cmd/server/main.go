package main

import (
	"fmt"
	"net/http"
)

// helloHandler handles requests to the root path.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Refactored World!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is healthy!")
}

// initializeRoutes sets up the routes for the server.
func initializeRoutes() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
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
