package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	unauthenticatedRouter := NewRouter()
	unauthenticatedRouter.GET("/", HandleHome)

	authenticatedRouter := NewRouter()
	authenticatedRouter.GET("/images/new", HandleImageNew)

	middleware := Middleware{}
	middleware.Add(unauthenticatedRouter)
	middleware.Add(http.HandlerFunc(AuthenticateRequest))
	middleware.Add(authenticatedRouter)

	http.ListenAndServe(":3000", middleware)
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
