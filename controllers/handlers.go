package controllers

import (
    "BookSearchGo/parsers"
    "html/template"
    "log"
    "net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("static/index.html")
    if err != nil {
        http.Error(w, "Could not serve page", http.StatusNotFound)
    } else {
        if err := t.Execute(w, nil); err != nil {
            log.Printf("Could not execute template: %s", err.Error())
        }
    }
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        http.Redirect(w, r, "", http.StatusPermanentRedirect)
    } else {
        _ = r.ParseForm()
        ResultData, err := parsers.ParseBookQuery(r.FormValue("query"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        t, err := template.ParseFiles("static/result.html")
        if err != nil {
            log.Fatal("Error: %s", err.Error())
        }
        if err := t.Execute(w, ResultData); err != nil {
            log.Printf(err.Error())
        }
    }
}
