package db

type Article struct {
	ID        uint     `json:"id"`
	Summary   string   `json:"summary"`
	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"-"`
}

func (a *Article) SetSummaryFromBody(body []byte) {

}

func (a *Article) SetImagesFromBody(body []byte) {

}

func (a *Article) SetThumbnail() {

}
