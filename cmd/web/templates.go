package main

import (
	"html/template"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/mihailtudos/portfolio/ui"
)

type templateData struct {
	Data        map[string]any
	CurrentYear int
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
		Data:        make(map[string]any),
	}
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Use fs.Glob() to get a slice of all filepaths in the ui.Files embedded
	// filesystem which match the pattern 'html/pages/**/*.go.html'. This essentially
	// gives us a slice of all the 'page' templates for the application
	pages, err := fs.Glob(ui.Files, "html/pages/**/*.go.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := strings.TrimPrefix(page, "html/pages/")

		// Create a slice containing the filepath patterns for the templates we parse.
		patterns := []string{
			"html/layouts/main.go.html",
			"html/partials/*.go.html",
			page,
		}

		// Use ParseFS() instead of ParseFiles() to parse the template files
		// from the ui.Files embedded filesystem.
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Return the map.
	return cache, nil
}
