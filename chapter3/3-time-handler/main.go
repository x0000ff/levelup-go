package main

import (
	"fmt"
	"net/http"
	"time"
)

// UptimeHandler writes the number of seconds since starting the response
type UptimeHandler struct {
	Started time.Time
}

// NewUptimeHandler ...
func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w,
		fmt.Sprintf("Current Uptime: %s", time.Since(h.Started)),
	)
}

func main() {
	http.Handle("/", NewUptimeHandler())
	http.ListenAndServe(":3000", nil)
}
