package controllers

import "net/http"

// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created successfully"))

}

// DeleteUser handles the deletion of a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User deleted successfully"))
}