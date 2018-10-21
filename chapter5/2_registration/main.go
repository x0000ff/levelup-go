package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := NewRouter()

	router.Handle("GET", "/", HandleHome)
	router.Handle("GET", "/register", HandleUserNew)

	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	middleware := Middleware{}
	middleware.Add(router)

	log.Fatal(http.ListenAndServe(":3000", middleware))
}

// NotFound ...
type NotFound struct{}

func (n *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// NewRouter ...
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	notFound := new(NotFound)
	router.NotFound = notFound
	return router
}
