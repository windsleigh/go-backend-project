package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler handles requests to the root path.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// HealthHandler handles requests to the /health path.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is healthy!")
}
