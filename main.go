package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// Handler function to serve SVG icons dynamically
func Handler(w http.ResponseWriter, r *http.Request) {
	// Get the requested icon name from the URL path
	iconName := filepath.Base(r.URL.Path)

	// Set the content type for SVG
	w.Header().Set("Content-Type", "image/svg+xml")

	// Build the full path to the requested icon file
	iconPath := filepath.Join("icons", iconName)

	// Serve the SVG icon if it exists
	if _, err := os.Stat(iconPath); err == nil {
		http.ServeFile(w, r, iconPath)
	} else {
		// Respond with 404 if the icon is not found
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Icon not found!")
	}
}

// Main function to start the server
func main() {
	http.HandleFunc("/", Handler)
	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
