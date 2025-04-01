package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Icon struct {
	Filename string
	Name     string
	Size     int
}

type PageData struct {
	Icons       []Icon
	SearchTerm  string
	CurrentPage int
	TotalPages  int
	PageSize    int
}

// Template functions
var templateFuncs = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"iterate": func(start, end int) []int {
		var result []int
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
		return result
	},
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
	// Parse the template file with custom functions
	tmpl, err := template.New("index.html").Funcs(templateFuncs).ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get the search query from the URL (e.g., ?search=icon-name)
	searchQuery := r.URL.Query().Get("search")

	// Get the page number from the URL, default to 1
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	// Set the number of icons per page
	pageSize := 105

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

	// Calculate total pages
	totalPages := (len(icons) + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	// Ensure page is not greater than total pages
	if page > totalPages {
		page = totalPages
	}

	// Calculate start and end indices for the current page
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(icons) {
		endIndex = len(icons)
	}

	// Get icons for the current page
	pageIcons := icons[startIndex:endIndex]

	// Render the template with the icons and pagination data
	err = tmpl.Execute(w, PageData{
		Icons:       pageIcons,
		SearchTerm:  searchQuery,
		CurrentPage: page,
		TotalPages:  totalPages,
		PageSize:    pageSize,
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
