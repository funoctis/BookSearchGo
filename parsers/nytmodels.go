package parsers

type NYTResp struct {
	Results []Book
}

type Book struct {
	PublishedDate string   `json:"published_date"`
	BookDetails   []Detail `json:"book_details"`
}

type Detail struct {
	Title         string
	Description   string
	Author        string
	PrimaryIsbn10 string `json:"primary_isbn10"`
}
