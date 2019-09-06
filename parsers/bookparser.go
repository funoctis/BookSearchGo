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

//ResultData is the return type for ParseBookQuery()
type ResultData struct {
	Title   string
	Volumes []Volume
}

//ParseBookQuery takes the query string entered by the user and
//fetches a response from Google Books API using BOOKSAPIKEY.
//Returns list of volumes along with some info about them.
func ParseBookQuery(query string) (*ResultData, error) {

	baseUrl := "https://www.googleapis.com/books/v1/volumes?"
	escapedQuery := strings.ReplaceAll(query, " ", "+")
	searchQuery := fmt.Sprintf("q=%s", escapedQuery)
	filters := "fields=items(id,volumeInfo,accessInfo)"

	key := os.Getenv("BOOKSAPIKEY")
	if key == "" {
		err := fmt.Errorf("couldn't fetch Google Books API Key")
		log.Printf("ERROR while fetching response: %s", err.Error())
		return &ResultData{}, err
	}

	link := fmt.Sprintf("%s%s&%s&%s", baseUrl, searchQuery, filters, key)

	response, err := http.Get(link)
	if err != nil {
		log.Printf("ERROR while fetching response: %s", err.Error())
		return &ResultData{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("ERROR while reading response: %s", err.Error())
		return &ResultData{}, err
	}

	var resp Resp

	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		log.Printf("ERROR while unmarshaling response: %s", err.Error())
		return &ResultData{}, err
	}

	return &ResultData{Title: query, Volumes: resp.Items[:10]}, nil
}
