package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Exported Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to serve HTML or SVGs
	w.Header().Set("Content-Type", "image/svg+xml")

	// Serve different SVG icons based on the URL path
	switch r.URL.Path {
	case "/icon1":
		fmt.Fprintf(w, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path d="M12 0l3 12h-6l3-12zm3 18v6h-6v-6h6zm9-6h-6v6h6v-6zm-12 0h-6v6h6v-6zm-9 0h6v6h-6v-6z"/></svg>`)
	case "/icon2":
		fmt.Fprintf(w, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path d="M2 0v24h20v-24h-20zm16 2v20h-12v-20h12zm-3 8h-6v2h6v-2z"/></svg>`)
	default:
		// Default response for any other route
		w.WriteHeader(http.StatusNotFound) // Respond with a 404 for unknown routes
		fmt.Fprintf(w, "Icon not found!")
	}
}

type Icon struct {
	Filename string
	Name     string
	Size     int
}

func handler() {
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
