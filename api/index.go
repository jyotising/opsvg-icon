package handler

import (
	"fmt"
	"net/http"
)

// Handler function for the HTTP request
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// Entry point for Vercel
func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil) // This line is ignored by Vercel
}
