package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleUserNew ...
func HandleUserNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Display Home Page
	http.Redirect(w, r, "/users/new", http.StatusFound)
}
