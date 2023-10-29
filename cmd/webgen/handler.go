package main

import (
	"WebPrac/pkg/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (pl *ProcessingLog) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pl.notFound(w)
		return
	}
	a, err := pl.article.LastUs()
	if err != nil {
		pl.serverError(w, err)
		return
	}
	//for _, artilc := range a {
	//	fmt.Fprintf(w, "%v\n", artilc)
	//}

	data := &tmData{Articles: a}

	files := []string{"./iu/html/home.p.tmpl"}
	tm, err := template.ParseFiles(files...)
	if err != nil {
		pl.serverError(w, err)
		return
	}
	err = tm.Execute(w, data)
	if err != nil {
		pl.serverError(w, err)
	}
}

func (pl *ProcessingLog) ShowArticles(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		pl.notFound(w)
		return
	}

	a, err := pl.article.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNotReport) {
			pl.notFound(w)
		} else {
			pl.serverError(w, err)
		}
		return
	}

	data := &tmData{Article: a}

	files := []string{"./iu/html/visual.p.tmpl"}
	tm, err := template.ParseFiles(files...)
	if err != nil {
		pl.serverError(w, err)
		return
	}
	err = tm.Execute(w, data)
	if err != nil {
		pl.serverError(w, err)
	}

}

func (pl *ProcessingLog) CreatArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		files := []string{"./iu/html/create.p.tmpl"}
		tm, err := template.ParseFiles(files...)
		if err != nil {
			pl.serverError(w, err)
			return
		}
		//tm.ExecuteTemplate(w, "create", nil)
		err = tm.Execute(w, nil)
		if err != nil {
			pl.serverError(w, err)
		}
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	//expires := r.FormValue("expires")

	//if r.Method != http.MethodPost {
	//	w.Header().Set("Allow", http.MethodPost)
	//	pl.clientError(w, http.StatusMethodNotAllowed)
	//	return
	//}
	//title := "..."
	//content := "..."
	expires := "7"
	//
	id, err := pl.article.Insert(title, content, expires)
	if err != nil {
		pl.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/articles?id=%d", id), http.StatusSeeOther)
}
