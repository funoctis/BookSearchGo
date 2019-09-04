package parsers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

var resp Resp

type ResultData struct {
    Title string
    Volumes []Volume
}

func ParseBookQuery(query string) ResultData {
    baseUrl := "https://www.googleapis.com/books/v1/volumes?"
    query = strings.ReplaceAll(query, " ", "+")
    searchQuery := fmt.Sprintf("q=%s", query)
    filters := "download=epub&fields=items(id,volumeInfo,accessInfo)"
    key:= os.Getenv("BOOKSAPIKEY")
    if key == "" {
        panic("Key not loaded")
    }
    link := baseUrl + searchQuery + "&" + filters + "&" + key

    response, err := http.Get(link)
    if err != nil {
        panic(err.Error())
    }
    var responseData, _ = ioutil.ReadAll(response.Body)

    err = json.Unmarshal(responseData, &resp)
    if err != nil {
        panic(err.Error())
    }

    return ResultData{Title:query, Volumes: resp.Items[:10]}
}
