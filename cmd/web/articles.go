package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mihailtudos/portfolio/internal/data"
	"github.com/mihailtudos/portfolio/internal/validator"
)

type articleCreateForm struct {
	Title   string
	Content string
	validator.Validator
}

func (app *application) listArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := app.Models.Articles.GetAll()
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Data["articles"] = articles

	app.render(w, r, http.StatusOK, "articles/index.go.html", data)
}

func (app *application) createArticle(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = articleCreateForm{}

	fmt.Printf("\n\n\nrestet form data:  %+v \n\n\n", data.Form)

	app.render(w, r, http.StatusOK, "articles/create.go.html", data)
}

func (app *application) storeArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := articleCreateForm{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This filed is required")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")

	fmt.Printf("%+v", form.FieldErrors)

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusOK, "articles/create.go.html", data)
		return
	}

	// Get the file from the form
	if _, ok := r.MultipartForm.File["image"]; ok {
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Unable to get file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		filePath, err := app.Services.Assets.SaveFile(file, handler, "./ui/static/img/articles/")
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		fmt.Println(filePath)
	}

	a := data.Article{
		Title:   form.Title,
		Content: form.Content,
	}

	err = app.Models.Articles.Insert(&a)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/articles/show/%d", a.ID), http.StatusSeeOther)
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
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Data["article"] = article

	app.render(w, r, http.StatusOK, "articles/show.go.html", data)
}
