package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HelloHandler handles requests to the root path.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // Get the 'name' query parameter
	if name == "" {
		name = "World" // Default to "World" if no name is provided
	}

	response := map[string]string{"greeting": fmt.Sprintf("Hello, %s!", name)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LoggingMiddleware logs incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler
	})
}
