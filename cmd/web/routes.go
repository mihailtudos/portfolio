package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mihailtudos/portfolio/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/projects", app.projects)
	router.HandlerFunc(http.MethodGet, "/resume", app.resume)

	router.HandlerFunc(http.MethodGet, "/articles", app.listArticles)
	router.HandlerFunc(http.MethodPost, "/articles", app.storeArticle)
	router.HandlerFunc(http.MethodGet, "/articles/:id", app.showArticle)
	fileServer := http.FileServer(http.FS(ui.Files))

	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	return router
}
