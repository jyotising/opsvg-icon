package main

import (
	"fmt"
	"net/http"
)

// The HTTP handler for this function
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func main() {
	http.HandleFunc("/api/hello", Handler) // Use the specific route for deployment
	// Start the server on port 8080 (for local testing)
	http.ListenAndServe(":8080", nil)
}
