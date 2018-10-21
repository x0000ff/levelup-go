package main

// AuthenticateRequest ...
import "net/http"

// AuthenticateRequest ...
func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	// Redirect the user to login if he's not authenticated
	authenticated := false
	if !authenticated {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}
