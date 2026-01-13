package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// userStore simulates an in-memory database.
// Key: username, Value: hashed password
var userStore = make(map[string]string)

// reader is a shared input reader for the CLI
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("=== Secure Login System v1 ===")

	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := readInput()

		switch choice {
		case "1":
			handleRegister()
		case "2":
			handleLogin()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// handleRegister manages the user registration flow
func handleRegister() {
	fmt.Println("\n--- Registration ---")

	fmt.Print("Enter Username: ")
	username, _ := readInput()
	if username == "" {
		fmt.Println("Error: Username cannot be empty.")
		return
	}

	// Check if user already exists
	if _, exists := userStore[username]; exists {
		fmt.Println("Error: User already exists.")
		return
	}

	fmt.Print("Enter Password: ")
	password, _ := readInput()
	if password == "" {
		fmt.Println("Error: Password cannot be empty.")
		return
	}

	// Hash the password using bcrypt
	// MinCost is used for testing speed, DefaultCost (10) is standard for production
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", err)
		return
	}

	// Store in memory
	userStore[username] = string(hashedPassword)
	fmt.Println("Registration Successful!")
}

// handleLogin manages the authentication and MFA flow
func handleLogin() {
	fmt.Println("\n--- Login ---")

	fmt.Print("Enter Username: ")
	username, _ := readInput()

	fmt.Print("Enter Password: ")
	password, _ := readInput()

	// 1. Fetch stored hash
	storedHash, exists := userStore[username]
	if !exists {
		// Generic error message prevents username enumeration attacks
		fmt.Println("Error: Invalid Credentials")
		return
	}

	// 2. Verify Password
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		fmt.Println("Error: Invalid Credentials")
		return
	}

	// 3. MFA Layer
	fmt.Println("Password Verified. Initiating MFA...")

	otp, err := generateSecureOTP()
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return
	}

	// Simulation of SMS
	fmt.Printf("\n[SMS SIMULATION] Your OTP is: %s\n\n", otp)

	fmt.Print("Enter OTP: ")
	inputOTP, _ := readInput()

	if inputOTP == otp {
		fmt.Println("\n>>> Login Successful! Welcome,", username, "<<<")
	} else {
		fmt.Println("\nError: Invalid OTP. Access Denied.")
	}
}

// generateSecureOTP generates a cryptographically secure 6-digit number
func generateSecureOTP() (string, error) {
	// 900000 + 100000 ensures we always get a 6-digit number (100000 to 999999)
	max := big.NewInt(900000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	// Add 100000 to ensure no leading zeros issues (e.g., getting "012345")
	return fmt.Sprintf("%d", n.Int64()+100000), nil
}

// readInput is a helper to read and trim strings from Stdin
func readInput() (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
