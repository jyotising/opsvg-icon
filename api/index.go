package main

import (
	"fmt"
	"net/http"
)

// Handler is the function that Vercel will use to handle incoming HTTP requests.
// Make sure it is capitalized (exported) so that Vercel can find it.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/html, text/plain, or image/svg+xml based on your needs
	w.Header().Set("Content-Type", "text/html")

	// Write a simple response (you can customize this to serve your icons or content)
	fmt.Fprintf(w, "<h1>Hello, Welcome to the SVG Icon Library!</h1>")
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
