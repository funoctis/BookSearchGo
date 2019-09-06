package parsers

/*
Response schema  areas used by the application:
{
    "id": {
    "type": "string"
    },
    "volumeInfo": {
    "type": "object",
    "properties": {
        "title": {
            "type": "string"
        },
        "authors": {
            "type": "array",
            "items": [
                {
                    "type": "string"
                }
            ]
        },
        "publishedDate": {
            "type": "string"
        },
        "description": {
            "type": "string"
        },
        "imageLinks": {
            "type": "object",
            "properties": {
                "thumbnail": {
                    "type": "string"
                }
            }
        }
    }
}
*/

//VolumeInfo contains info for each volume(book)
type VolumeInfo struct {
	Title         string
	Authors       []string
	PublishedDate string
	Description   string
	InfoLink      string
	ImageLinks    ImageLinks
}

//ImageLinks is used by Bootstrap Media Object as the image source
type ImageLinks struct {
	Thumbnail string
}

//Volume is an individual book, multiple books are received as a JSON array
type Volume struct {
	Id         string
	VolumeInfo VolumeInfo
}

//Resp is the main struct for received JSON
type Resp struct {
	Items []Volume
}
