package controllers

import (
    "log"
    "net/http"
)

//LoginHandler is the handler function for the register page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "login", nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Could not serve page", http.StatusNotFound)
    }
}
