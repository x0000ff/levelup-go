package main

import "net/http"

// MiddlewareResponseWriter ...
type MiddlewareResponseWriter struct {
	http.ResponseWriter
	written bool
}

// NewMiddlewareResponseWriter ...
func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	return &MiddlewareResponseWriter{
		ResponseWriter: w,
	}
}

// Write ...
func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(bytes)
}

// WriteHeader ...
func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

// Middleware ...
type Middleware []http.Handler

// Add adds a handler to the Middleware
func (m *Middleware) Add(handler http.Handler) {
	*m = append(*m, handler)
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Write the supplied ResponseWriter
	mw := NewMiddlewareResponseWriter(w)

	// Loop through all the registered handlers

	for _, handler := range m {
		// Call the handler with out NewMiddlewareResponseWriter
		handler.ServeHTTP(mw, r)

		// If there was a write, stop processing
		if mw.written {
			return
		}
	}

	// If no handlers wrote to the response, it's a 404
	http.NotFound(w, r)
}
