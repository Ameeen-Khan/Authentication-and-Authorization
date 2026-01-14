package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TokenResponse represents tokens returned after auth code exchange
type TokenResponse struct {
	IDToken     string
	AccessToken string
}

// simulateRedirect simulates redirect to Identity Provider
func simulateRedirect() {
	fmt.Println("Redirecting to Identity Provider...")
}

// simulateLogin simulates IdP login and returns auth code
func simulateLogin() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	_, _ = reader.ReadString('\n')

	fmt.Print("Enter password: ")
	_, _ = reader.ReadString('\n')

	// Pretend authentication succeeded
	authCode := "XYZ123"
	fmt.Println("Auth Code:", authCode)

}

// exchangeAuthCode simulates auth code → token exchange
func exchangeAuthCode(inputCode string) (*TokenResponse, error) {
	if inputCode != "XYZ123" {
		return nil, fmt.Errorf("invalid authorization code")
	}

	return &TokenResponse{
		IDToken:     "abc.id.sig",
		AccessToken: "xyz.access.sig",
	}, nil
}

// verifyToken simulates token verification
func verifyToken(token string) bool {
	if token == "" {
		return false
	}

	// Very basic structure check: xxx.yyy.zzz
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}

	// Simulated claim check
	if parts[0] == "" || parts[1] == "" {
		return false
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Step 1: Redirect
	simulateRedirect()

	// Step 2: Login and get auth code
	simulateLogin()

	// Step 3: User enters auth code
	fmt.Print("Enter received auth code: ")
	inputCode, _ := reader.ReadString('\n')
	inputCode = strings.TrimSpace(inputCode)

	// Step 4: Exchange auth code for tokens
	tokens, err := exchangeAuthCode(inputCode)
	if err != nil {
		fmt.Println("Token Exchange Failed:", err)
		return
	}

	fmt.Println("Received Tokens:")
	fmt.Println("ID Token:", tokens.IDToken)
	fmt.Println("Access Token:", tokens.AccessToken)

	// Step 5: Verify tokens
	if verifyToken(tokens.IDToken) && verifyToken(tokens.AccessToken) {
		fmt.Println("Token Verified")
	} else {
		fmt.Println("Invalid Token")
	}
}
