package parsers

type ReadingModes struct {
    PageCount int
    AverageRating int
}

type VolumeInfo struct {
    Title string
    Authors []string
    PublishedDate string
    Description string
    InfoLink string
    ImageLinks ImageLinks
    ReadingModes ReadingModes
}

type ImageLinks struct {
    Thumbnail string
}

type Epub struct {
    IsAvailable bool
    DownloadLink string
}

type Pdf struct {
    IsAvailable bool
    DownloadLink string
}

type AccessInfo struct {
    Epub Epub
    Pdf Pdf
    WebReaderLink string
}

type Volume struct {
    Id string
    VolumeInfo VolumeInfo
    AccessInfo AccessInfo
}

type Resp struct {
    Items []Volume
}