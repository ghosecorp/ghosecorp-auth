package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users []User

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GET /users
func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// POST /users
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	u.ID = len(users) + 1
	users = append(users, u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/create", createUser)

	fmt.Println("Server running on http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}
