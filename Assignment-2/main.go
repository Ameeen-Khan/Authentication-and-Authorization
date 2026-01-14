package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
SECRET KEY
- In real systems: store in ENV or secret manager
- Never hardcode in production
*/
var jwtSecret = []byte("super-secret-key")

// CustomClaims defines our JWT payload
type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// generateToken creates a signed JWT
func generateToken(userID string) (string, error) {
	claims := CustomClaims{
		Role: "user",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// validateToken verifies signature, expiry, and decodes claims
func validateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Enforce signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func main() {
	// Step 1: Generate JWT
	token, err := generateToken("user-123")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated Token:")
	fmt.Println(token)

	fmt.Println("\n--- Validating Token ---")

	// Step 2: Validate JWT
	claims, err := validateToken(token)
	if err != nil {
		fmt.Println("Token Invalid:", err)
		return
	}

	fmt.Println("Token Valid")
	fmt.Println("User ID:", claims.Subject)
	fmt.Println("Role:", claims.Role)
	fmt.Println("Expires At:", claims.ExpiresAt.Time)
}
