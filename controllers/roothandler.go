package controllers

import (
	"html/template"
	"log"
	"net/http"
)

//Parsing and caching the templates beforehand, to be executed later.
var templates = template.Must(template.ParseGlob("static/*"))

//RootHandler is the handler function for index page
func RootHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not serve page", http.StatusNotFound)
	}
}
