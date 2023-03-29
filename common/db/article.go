package db

import "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"

const MAX_SUMMARY_LENGTH = 100

type Article struct {
	ID        uint     `json:"id"`
	Summary   string   `json:"summary"`
	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"-"`
}

func Save(article *Article) {
	db.Save(article)
}

// CreateArticle creates article from body. but doesn't save it
func CreateArticle(id uint, body []byte) *Article {
	article := new(Article)

	article.ID = id
	article.setData(body)

	return article
}

func DeleteByID(id uint) {
	db.Where("id = ?", id).Delete(&Article{}, id)
}

func GetImageURLsByID(id uint) []string {
	urls := []string{}
	db.Where("id = ?", id).Select("images").Take(urls)
	return urls
}

func (a *Article) setData(body []byte) {
	a.setSummaryFromBody(body)
	a.setImagesFromBody(body)
	a.setThumbnail()
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
