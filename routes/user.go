package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// User struct defines the structure of a user
type User struct {
	ID   int    `json:"id"`   // ID of the user
	Name string `json:"name"` // Name of the user
}

// Simulated in-memory database of users
var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Smith"},
}

// userHandler handles requests for the /user endpoint (GET and POST)
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Switch based on HTTP method (GET or POST)
	switch r.Method {
	case http.MethodGet:
		// Return the list of users as JSON
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		// Create a new user
		var newUser User
		// Decode the incoming JSON request body into the User struct
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Assign a new ID based on the current length of the users slice
		newUser.ID = len(users) + 1
		users = append(users, newUser) // Append the new user to the in-memory database

		// Respond with a 201 status (Created) and the new user data in JSON format
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// userDetailHandler handles requests for /user/{id} (PUT and DELETE)
func userDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the user ID from the URL (e.g., /user/1, /user/2, etc.)
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2]) // Convert the ID from string to int
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Switch based on HTTP method (PUT or DELETE)
	switch r.Method {
	case http.MethodPut:
		// Update an existing user
		var updatedUser User
		// Decode the incoming request body into the User struct
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Loop through the users and find the one with the matching ID
		for i, user := range users {
			if user.ID == id {
				users[i].Name = updatedUser.Name    // Update the user's name
				json.NewEncoder(w).Encode(users[i]) // Return the updated user
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)

	case http.MethodDelete:
		// Delete a user by ID
		for i, user := range users {
			if user.ID == id {
				// Remove the user from the slice
				users = append(users[:i], users[i+1:]...)
				w.WriteHeader(http.StatusNoContent) // Return 204 No Content (successful deletion)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Set up the HTTP routes
	http.HandleFunc("/user", userHandler)        // Handles GET & POST methods
	http.HandleFunc("/user/", userDetailHandler) // Handles PUT & DELETE methods for specific user IDs

	// Start the HTTP server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
