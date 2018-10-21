package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleUserNew ...
func HandleUserNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Display Home Page
	RenderTemplate(w, r, "users/new", nil)
}
