package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleHome ...
func HandleHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Display Home Page
	RenderTemplate(w, r, "index/home", nil)
}
