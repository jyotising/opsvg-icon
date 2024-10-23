package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message for Vercel
	fmt.Fprintf(w, "Hello from the Vercel deployment!")
}

func main() {
	// Register the handler function for the root URL "/"
	http.HandleFunc("/", handler)

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil)
}
