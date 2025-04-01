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
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get the search query from the URL (e.g., ?search=icon-name)
	searchQuery := r.URL.Query().Get("search")

	// Load all icons from the /icons folder
	icons, err := loadIcons()
	if err != nil {
		log.Printf("Error loading icons: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// If there's a search query, filter the icons
	if searchQuery != "" {
		icons = searchIcons(icons, searchQuery)
	}

	// Render the template with the icons and the search query
	err = tmpl.Execute(w, struct {
		Icons      []Icon
		SearchTerm string
	}{
		Icons:      icons,
		SearchTerm: searchQuery,
	})
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func loadIcons() ([]Icon, error) {
	var icons []Icon
	err := filepath.Walk("icons", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".svg" {
			icons = append(icons, Icon{
				Filename: filepath.Base(path),
				Name:     strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
				Size:     48, // Default size
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return icons, nil
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
