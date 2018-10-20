package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, "index/home", nil)
	})

	assetsHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/")))
	mux.Handle("/assets/", assetsHandler)

	// Super simple static webserver
	log.Fatal(http.ListenAndServe(":3000", mux))
}
