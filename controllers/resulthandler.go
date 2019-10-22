package controllers

import (
    "BookSearchGo/parsers"
    "log"
    "net/http"
    "net/url"
)

//ResultHandler is the handler function for /result
//Parses and executes template for displaying the response data from parsers.ParseBookQuery()
//Redirects to index page if no POST request
func ResultHandler(w http.ResponseWriter, r *http.Request) {
    urlString, err := url.Parse(r.URL.String())
    if err != nil {
        http.Error(w, "ERROR while parsing url: "+err.Error(), http.StatusInternalServerError)
    }

    rawQuery := urlString.RawQuery

    if rawQuery == "" {
        http.Redirect(w, r, "", http.StatusPermanentRedirect)
    } else {
        m, err := url.ParseQuery(rawQuery)
        query := m["query"][0]

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        ResultData, err := parsers.ParseBookQuery(query)
        if err != nil {
            log.Println(err.Error())
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        if err := templates.ExecuteTemplate(w, "result", ResultData); err != nil {
            log.Printf(err.Error())
            http.Error(w, "Could not serve page", http.StatusNotFound)
        }
    }
}
