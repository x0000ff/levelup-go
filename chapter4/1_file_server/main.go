package main

import (
	"log"
	"net/http"
)

func main() {
	assetsHandler := http.FileServer(http.Dir("assets/"))

	// Super simple static webserver
	log.Fatal(http.ListenAndServe(":3000", assetsHandler))
}
