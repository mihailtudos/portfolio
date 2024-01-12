package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := &application{
		templateCache: templateCache,
	}

	log.Println("starting server at port 8080...🔥")
	log.Fatal(http.ListenAndServe(":8080", app.routes()))
}
