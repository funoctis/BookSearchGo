package controllers

import (
    "log"
    "net/http"
)

//RegisterHandler is the handler function for the register page
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "register", nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Could not serve page", http.StatusNotFound)
    }
}
