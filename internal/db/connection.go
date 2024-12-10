package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect initializes the database connection
func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %s", err)
		return nil, err
	}

	log.Println("Database connected successfully!")
	return db, nil
}

// InitializeDatabase runs an initialization SQL script
func InitializeDatabase(db *sql.DB) error {
	script, err := ioutil.ReadFile("scripts/init.sql")
	if err != nil {
		return fmt.Errorf("failed to read init.sql: %w", err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		return fmt.Errorf("failed to execute init.sql: %w", err)
	}

	log.Println("Database initialized successfully!")
	return nil
}
