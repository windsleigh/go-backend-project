package main

import (
	"go-backend-project/internal/db"
	"go-backend-project/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func initializeRoutes() {
	http.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HelloHandler)))
	http.Handle("/health", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HealthHandler)))
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running at http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}
	defer dbConn.Close()

	// Run the SQL initialization script
	if err := db.InitializeDatabase(dbConn); err != nil {
		log.Fatalf("Failed to initialize the database: %s", err)
	}

	handlers.SetDatabase(dbConn)

	initializeRoutes()
	startServer()
}
