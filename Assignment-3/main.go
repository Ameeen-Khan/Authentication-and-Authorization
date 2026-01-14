package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Document represents a protected resource
type Document struct {
	ID      int
	Owner   string
	Content string
}

// In-memory document store
var documents = []Document{
	{ID: 1, Owner: "userA", Content: "A's secret"},
	{ID: 2, Owner: "userB", Content: "B's secret"},
}

// fetchDocument finds a document by ID
func fetchDocument(id int) (*Document, error) {
	for _, doc := range documents {
		if doc.ID == id {
			return &doc, nil
		}
	}
	return nil, fmt.Errorf("document not found")
}

// canAccess checks authorization rules
func canAccess(username, role string, doc *Document) bool {
	if role == "admin" {
		return true
	}
	return doc.Owner == username
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Input: username
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Input: role
	fmt.Print("Enter role (user/admin): ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	// Input: document ID
	fmt.Print("Enter document ID: ")
	docIDInput, _ := reader.ReadString('\n')
	docIDInput = strings.TrimSpace(docIDInput)

	docID, err := strconv.Atoi(docIDInput)
	if err != nil {
		fmt.Println("Invalid document ID")
		return
	}

	// Fetch document
	doc, err := fetchDocument(docID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Authorization check
	if canAccess(username, role, doc) {
		fmt.Println("Access Granted")
		fmt.Println("Document Content:", doc.Content)
	} else {
		fmt.Println("Access Denied")
	}
}
