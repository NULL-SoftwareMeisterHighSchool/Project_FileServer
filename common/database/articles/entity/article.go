package article_entity

import (
	"strings"
	"time"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

const MAX_SUMMARY_LENGTH = 400

// should change those urls into url.URL type
type Article struct {
	ID        uint        `gorm:"autoIncrement"`
	AuthorID  uint        `gorm:"not null"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
	Title     string      `gorm:"type:varchar(2048),index:,class:FULLTEXT,option:WITH PARSER ngram"`
	Summary   string      `gorm:"type:varchar(400)"`
	Thumbnail string      `gorm:"type:varchar(2048)"`
	Images    string      `gorm:"type:text"`
	Body      ArticleBody `gorm:"constraint:OnDelete:CASCADE"`
	Views     uint64
}

type ArticleBody struct {
	ArticleID uint
	Text      []byte `gorm:"type:longtext,index:,class:FULLTEXT,option:WITH PARSER ngram"`
}

func New() *Article {
	return &Article{}
}

func (a *Article) SetID(articleID uint) *Article {
	a.ID = articleID
	return a
}

func (a *Article) SetAuthorID(authorID uint) *Article {
	a.AuthorID = authorID
	return a
}

func (a *Article) SetBody(body []byte) *Article {
	a.Body = ArticleBody{
		ArticleID: a.ID,
		Text:      body,
	}
	return a
}

func (a *Article) SetTitle(title string) *Article {
	a.Title = title
	return a
}

func (a *Article) SetSummary(body []byte) *Article {
	plainText := util.SanitizeExceptPlainText(body)
	minLength := util.GetMin(len(plainText), MAX_SUMMARY_LENGTH)
	a.Summary = string(plainText[:minLength])
	return a
}

func (a *Article) SetImagesAndThumbnail(images []string) *Article {
	a.Images = strings.Join(images, "^")
	a.setThumbnail()
	return a
}

func (a *Article) setThumbnail() {
	if len(a.Images) == 0 {
		a.Thumbnail = ""
		return
	}
	a.Thumbnail = strings.Split(a.Images, "^")[0]
}
