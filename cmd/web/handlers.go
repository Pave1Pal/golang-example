package main

import (
	"net/http"
	"strconv"
	"text/template"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFoundError(w)
		return
	}

	ts, err := template.ParseFiles("../../ui/html/home-page.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Println("Ok")
}

func (app *Application) showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientError(w, 405)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.serverError(w, err)
		return
	}
	w.Write([]byte(strconv.Itoa(id)))
}

func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, 405)
		return
	}
	w.Write([]byte("form for creating new snippet"))
}
