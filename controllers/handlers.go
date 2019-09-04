package controllers

import (
    "BookSearchGo/parsers"
    "html/template"
    "log"
    "net/http"
)

var templates = template.Must(template.ParseGlob("static/*"))

func RootHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index", nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Could not serve page", http.StatusNotFound)
    }
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        http.Redirect(w, r, "", http.StatusPermanentRedirect)
    } else {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        ResultData, err := parsers.ParseBookQuery(r.FormValue("query"))
        if err != nil {
            log.Println(err.Error())
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        if err := templates.ExecuteTemplate(w, "result",ResultData); err != nil {
            log.Printf(err.Error())
            http.Error(w, "Could not serve page", http.StatusNotFound)
        }
    }
}
