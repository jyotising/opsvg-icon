package handler

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var templates *template.Template

func init() {
	// Load HTML templates
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

// Vercel handler for Go
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve static files or templates based on URL
		serveFileOrTemplate(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Serve static files or HTML templates
func serveFileOrTemplate(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		// Render the main HTML page
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		// Serve static files (CSS, JS, SVG, etc.)
		staticPath := filepath.Join("static", filepath.Clean(r.URL.Path))
		if _, err := os.Stat(staticPath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, staticPath)
	}
}
