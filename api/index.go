package main

import (
	"fmt"
	"net/http"
)

// The HTTP handler for this function
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// The entry point for Vercel
func main() {
	http.HandleFunc("/", Handler) // This is required for local development; Vercel uses individual handlers
	// No need to run ListenAndServe; Vercel handles this
}
