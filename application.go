package main

import (
	"BookSearchGo/controllers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

//Creates a web server which listens and servers on specified PORT
func main() {
	var port string

	err := godotenv.Load()
	if err != nil {
		log.Printf("ERROR could not load from .env file: %s", err.Error())
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", controllers.RootHandler)
	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/welcome", controllers.WelcomeHandler)
	http.HandleFunc("/result", controllers.ResultHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
