package db

import "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"

const MAX_SUMMARY_LENGTH = 100

type Article struct {
	ID        uint     `json:"id"`
	Summary   string   `json:"summary"`
	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"-"`
}

func CreateArticleFromBody(body []byte) *Article {
	article := new(Article)

	article.setSummaryFromBody(body)
	article.setImagesFromBody(body)
	article.setThumbnail()

	db.Create(article)

	return article
}

func (a *Article) setSummaryFromBody(body []byte) {
	plainText := util.SanitizeExceptPlainText(body)
	minLength := util.GetMin(len(plainText), MAX_SUMMARY_LENGTH)
	a.Summary = string(plainText[:minLength])
}

func (a *Article) setImagesFromBody(body []byte) {
	a.Images = util.ExtractImageURL(body)
}

func (a *Article) setThumbnail() {
	if len(a.Images) == 0 {
		a.Thumbnail = ""
	} else {
		a.Thumbnail = a.Images[0]
	}
}
