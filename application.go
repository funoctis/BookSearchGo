package main

import (
    "BookSearchGo/controllers"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"
)

func main() {
    var port string
    _ = godotenv.Load()
    port = os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    http.HandleFunc("/", controllers.RootHandler)
    http.HandleFunc("/result", controllers.ResultHandler)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
