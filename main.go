package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Icon represents an SVG icon with its filename and name
type Icon struct {
	Filename string
	Name     string
	Size     int
}

// LoadIcons loads SVG files from the icons directory
func LoadIcons() ([]Icon, error) {
	var icons []Icon
	iconsDir := "./icons"

	err := filepath.Walk(iconsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only consider .svg files
		if !info.IsDir() && filepath.Ext(path) == ".svg" {
			icons = append(icons, Icon{
				Filename: info.Name(),
				Name:     info.Name(),
				Size:     100, // Set a default size (you can adjust this)
			})
		}
		return nil
	})

	return icons, err
}

// ServeIcons serves the HTML with icons
func ServeIcons(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	icons, err := LoadIcons()
	if err != nil {
		http.Error(w, "Error loading icons", http.StatusInternalServerError)
		return
	}

	// Pass icons to the template
	tmpl.Execute(w, struct {
		Icons []Icon
	}{
		Icons: icons,
	})
}

func main() {
	// Serve static files (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve icon files (SVG)
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))

	// Serve HTML template
	http.HandleFunc("/", ServeIcons)

	// Get the port from environment variable or use 8080 as default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)

	// Start the HTTP server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
