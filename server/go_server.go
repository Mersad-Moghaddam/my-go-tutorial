package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct to represent a User
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handler for the welcome page
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our simple web server!")
}

// Handler for the user info API
func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "John Doe", Age: 30}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Define routes
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/user", userHandler)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
