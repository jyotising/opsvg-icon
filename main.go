package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Handler is the function Vercel uses to handle HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML, SVG, or any other type as needed
	w.Header().Set("Content-Type", "text/html")

	// This is a basic response to test that your handler is working
	fmt.Fprintf(w, "<h1>Welcome to the OpSVG Icon Library!</h1>")
}

type Icon struct {
	Filename string
	Name     string
	Size     int
}

func main() {
	// Serve static files from /static and /icons folders
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))

	// Handle requests for the home page
	http.HandleFunc("/", homeHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the template file
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Get the search query from the URL (e.g., ?search=icon-name)
	searchQuery := r.URL.Query().Get("search")

	// Load all icons from the /icons folder
	icons := loadIcons()

	// If there's a search query, filter the icons
	if searchQuery != "" {
		icons = searchIcons(icons, searchQuery)
	}

	// Render the template with the icons and the search query
	tmpl.Execute(w, struct {
		Icons      []Icon
		SearchTerm string
	}{
		Icons:      icons,
		SearchTerm: searchQuery, // Pass the search term to the template
	})
}

func loadIcons() []Icon {
	var icons []Icon
	filepath.Walk("icons", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".svg" {
			icons = append(icons, Icon{
				Filename: filepath.Base(path),
				Name:     strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
				Size:     100, // Default size
			})
		}
		return nil
	})
	return icons
}

func searchIcons(icons []Icon, query string) []Icon {
	var filteredIcons []Icon
	query = strings.ToLower(query) // Make search case-insensitive
	for _, icon := range icons {
		if strings.Contains(strings.ToLower(icon.Name), query) {
			filteredIcons = append(filteredIcons, icon)
		}
	}
	return filteredIcons
}
