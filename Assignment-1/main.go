package main

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users map[string]string
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: make(map[string]string),
	}
}

func (a *AuthService) Register(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password required")
	}

	if _, exists := a.users[username]; exists {
		return errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.users[username] = string(hash)
	return nil
}

func (a *AuthService) Authenticate(username, password string) error {
	hash, exists := a.users[username]
	if !exists {
		return errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}

func generateOTP() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()+100000), nil
}

// ---------------- CLI ----------------

var reader = bufio.NewReader(os.Stdin)

func main() {
	auth := NewAuthService()

	fmt.Println("=== Secure Login System ===")

	for {
		fmt.Println("\n1. Register\n2. Login\n3. Exit")
		fmt.Print("Choice: ")

		choice, err := readInput()
		if err != nil {
			fmt.Println("Input error")
			continue
		}

		switch choice {
		case "1":
			handleRegister(auth)
		case "2":
			handleLogin(auth)
		case "3":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func handleRegister(auth *AuthService) {
	fmt.Print("Username: ")
	u, _ := readInput()

	fmt.Print("Password: ")
	p, _ := readInput()

	if err := auth.Register(u, p); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Registration successful")
}

func handleLogin(auth *AuthService) {
	fmt.Print("Username: ")
	u, _ := readInput()

	fmt.Print("Password: ")
	p, _ := readInput()

	if err := auth.Authenticate(u, p); err != nil {
		fmt.Println("Error: Invalid credentials")
		return
	}

	otp, err := generateOTP()
	if err != nil {
		fmt.Println("OTP error")
		return
	}

	fmt.Println("[SMS] OTP:", otp)

	fmt.Print("Enter OTP: ")
	input, _ := readInput()

	if input != otp {
		fmt.Println("Invalid OTP")
		return
	}

	fmt.Println("Login successful!")
}

func readInput() (string, error) {
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}
