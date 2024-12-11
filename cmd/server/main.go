package main

import (
	"go-backend-project/internal/db"
	"go-backend-project/internal/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func initializeRoutes() {
	http.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HelloHandler)))
	http.Handle("/health", handlers.LoggingMiddleware(http.HandlerFunc(handlers.HealthHandler)))
	http.Handle("/login", handlers.LoggingMiddleware(http.HandlerFunc(handlers.LoginHandler)))

	// Protected routes
	http.Handle("/users", handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(handlers.GetUsersHandler))))
	http.Handle("/users/create", handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(handlers.CreateUserHandler))))
	http.Handle("/users/update", handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(handlers.UpdateUserHandler))))
	http.Handle("/users/delete", handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(handlers.DeleteUserHandler))))
}

func startServer() {
	port := ":8080"

	// Wrap all routes with CORS middleware
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(http.DefaultServeMux)

	log.Printf("Server is running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
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
