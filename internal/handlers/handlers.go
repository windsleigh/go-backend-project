package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var database *sql.DB // Global variable to hold the database connection

// SetDatabase sets the database connection for handlers to use.
func SetDatabase(dbConn *sql.DB) {
	database = dbConn
}

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

// HealthHandler handles requests to check the server's health.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LoggingMiddleware logs incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logLevel := os.Getenv("LOG_LEVEL")

		// Log only if log level is INFO or DEBUG
		if logLevel == "INFO" || logLevel == "DEBUG" {
			log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		}

		// Log additional details for DEBUG level
		if logLevel == "DEBUG" {
			log.Printf("Headers: %v", r.Header)
		}

		next.ServeHTTP(w, r) // Call the next handler
	})
}
