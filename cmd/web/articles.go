package main

import (
	"fmt"
	"net/http"

	"github.com/mihailtudos/portfolio/internal/data"
)

func (app *application) storeArticle(w http.ResponseWriter, r *http.Request) {
	a := data.Article{Title: "title new", Content: "content new"}

	app.Models.Articles.Insert(&a)
	fmt.Fprintf(w, "%+v", a)
}
