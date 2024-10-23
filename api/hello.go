package main

import (
	"fmt"
	"net/http"
)

// The HTTP handler for this function
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// This function serves as the entry point for Vercel
func main() {
	http.HandleFunc("/", Handler) // This line is not needed for Vercel, but fine for local testing
}
