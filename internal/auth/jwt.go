package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key") // Replace with a secure key

// Claims represents the JWT payload
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a given username
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates a given JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
