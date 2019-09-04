package parsers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
)

type ResultData struct {
    Title string
    Volumes []Volume
}

func ParseBookQuery(query string) ResultData {
    baseUrl := "https://www.googleapis.com/books/v1/volumes?"
    query = strings.ReplaceAll(query, " ", "+")
    searchQuery := fmt.Sprintf("q=%s", query)
    filters := "fields=items(id,volumeInfo,accessInfo)"
    key:= os.Getenv("BOOKSAPIKEY")
    if key == "" {
        log.Fatal("Books API Key not loaded")
    }

    link := baseUrl + searchQuery + "&" + filters + "&" + key
    response, err := http.Get(link)
    if err != nil {
        log.Fatal("Error: ", err.Error())
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Printf("Error while reading response: %s", err.Error())
    }
    var resp Resp
    err = json.Unmarshal(responseData, &resp)
    if err != nil {
        log.Printf("Error while unmarshaling json: %s", err.Error())
    }

    return ResultData{Title:query, Volumes: resp.Items[:10]}
}
