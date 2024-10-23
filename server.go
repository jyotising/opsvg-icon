package main

import (
	"fmt"
	"net/http"
)

// This function handles requests to the root URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the Vercel deployment!")
}

// Entry point for Vercel
func main() {
	http.HandleFunc("/", handler)     // Register the handler for the root URL
	http.ListenAndServe(":8080", nil) // Start the server
}
