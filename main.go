package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Icon struct {
	Filename string
	Name     string
	Size     int
}

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve static files from /static and /icons folders
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))

	// Handle requests for the home page
	http.HandleFunc("/", homeHandler)

	log.Printf("Server started at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
				Size:     48, // Default size
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
