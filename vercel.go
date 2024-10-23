package main

import (
	"fmt"
	"net/http"
)

// Handler function to handle HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// Entry point for Vercel
func main() {
	// Register the handler for the root URL
	http.HandleFunc("/", Handler)

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil)
}
