package controllers

import (
    "BookSearchGo/parsers"
    "html/template"
    "net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("static/index.html")
    if err := t.Execute(w, nil); err != nil {
        panic(err.Error())
    }
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        http.Redirect(w, r, "", http.StatusPermanentRedirect)
    } else {
        _ = r.ParseForm()
        ResultData := parsers.ParseBookQuery(r.FormValue("query"))

        t, _ := template.ParseFiles("static/result.html")
        if err := t.Execute(w, ResultData); err != nil {
            panic(err.Error())
        }
    }
}
