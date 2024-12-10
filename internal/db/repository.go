package db

import (
	"database/sql"
	"errors"
	"log"
)

// User represents a user in the database.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		log.Printf("Error creating user: %s", err)
		return err
	}
	return nil
}

// GetUsers retrieves all users from the database.
func GetUsers(db *sql.DB) ([]User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error retrieving users: %s", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Printf("Error scanning row: %s", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser updates an existing user's name or email.
func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	result, err := db.Exec(query, name, email, id)
	if err != nil {
		log.Printf("Error updating user: %s", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no user found with the given ID")
	}
	return nil
}

// DeleteUser removes a user from the database by ID.
func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user: %s", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no user found with the given ID")
	}
	return nil
}
