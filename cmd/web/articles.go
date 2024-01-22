package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mihailtudos/portfolio/internal/data"
)

func (app *application) storeArticle(w http.ResponseWriter, r *http.Request) {
	a := data.Article{Title: "title new 3", Content: "content new 3"}

	app.Models.Articles.Insert(&a)

	http.Redirect(w, r, fmt.Sprintf("/articles/%d", a.ID), http.StatusSeeOther)
}

func (app *application) showArticle(w http.ResponseWriter, r *http.Request) {
	idParam := httprouter.ParamsFromContext(r.Context()).ByName("id")

	// TODO: add validation
	id, err := strconv.Atoi(idParam)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	article, err := app.Models.Articles.Get(int64(id))
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFound(w, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Data["article"] = article

	app.render(w, r, http.StatusOK, "articles/show.go.html", data)
}

func (app *application) listArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := app.Models.Articles.GetAll()
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFound(w, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Data["articles"] = articles

	app.render(w, r, http.StatusOK, "articles/index.go.html", data)
}
