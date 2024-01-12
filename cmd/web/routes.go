package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mihailtudos/portfolio/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.home)
	fileServer := http.FileServer(http.FS(ui.Files))

	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	return router
}
