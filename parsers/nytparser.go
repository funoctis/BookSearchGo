package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	NYTBaseUrl = "https://api.nytimes.com/svc/books/v3/lists.json?list-name=hardcover-fiction"
)

type NYTResult struct {
	Results []Book
	//ImageLinks ImageLinks
}

func ParseNYTResult() (*NYTResult, error) {
	NYTKey := os.Getenv("NYTAPIKEY")
	if NYTKey == "" {
		err := fmt.Errorf("couldn't fetch Google Books API Key")
		log.Println(err.Error())
	}

	link := fmt.Sprintf("%s&api-key=%s", NYTBaseUrl, NYTKey)
	response, err := http.Get(link)
	if err != nil {
		log.Printf("ERROR while fetching response: %s", err.Error())
		return &NYTResult{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("ERROR while reading response: %s", err.Error())
		return &NYTResult{}, err
	}

	var resp NYTResp

	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		log.Printf("ERROR while unmarshaling response: %s", err.Error())
		return &NYTResult{}, err
	}

	return &NYTResult{Results: resp.Results}, nil
}
