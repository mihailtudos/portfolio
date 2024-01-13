package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home/index.go.html", templateData{})
}

func (app *application) articles(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "articles/index.go.html", templateData{})
}

func (app *application) resume(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "resume/index.go.html", templateData{})
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "projects/index.go.html", templateData{})
}
