package main

import (
	"fmt"
	"net/http"
)

func vercelHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message
	fmt.Fprintf(w, "Hello from the Vercel deployment!")
}

func init() {
	// Register the handler function for the root URL "/"
	http.HandleFunc("/", vercelHandler)
}
