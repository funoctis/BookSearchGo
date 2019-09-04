package parsers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

type ResultData struct {
    Title string
    Volumes []Volume
}

func ParseBookQuery(query string) (ResultData, error) {
    baseUrl := "https://www.googleapis.com/books/v1/volumes?"
    escapedQuery := strings.ReplaceAll(query, " ", "+")
    searchQuery := fmt.Sprintf("q=%s", escapedQuery)
    filters := "fields=items(id,volumeInfo,accessInfo)"
    key := os.Getenv("BOOKSAPIKEY")
    if key == "" {
        return ResultData{}, fmt.Errorf("couldn't fetch Google Books API Key")
    }

    link := fmt.Sprintf("%s%s&%s&%s", baseUrl, searchQuery, filters, key)

    response, err := http.Get(link)
    if err != nil {
        return ResultData{}, err
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return ResultData{}, err
    }
    var resp Resp
    err = json.Unmarshal(responseData, &resp)
    if err != nil {
        return ResultData{}, err
    }

    return ResultData{Title: query, Volumes: resp.Items[:10]}, nil
}
