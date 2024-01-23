package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/mihailtudos/portfolio/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/projects", app.projects)
	router.HandlerFunc(http.MethodGet, "/resume", app.resume)

	router.HandlerFunc(http.MethodGet, "/articles", app.listArticles)
	router.HandlerFunc(http.MethodGet, "/articles/show/:id", app.showArticle)
	router.HandlerFunc(http.MethodPost, "/articles/create", app.storeArticle)
	router.HandlerFunc(http.MethodGet, "/articles/create", app.createArticle)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
