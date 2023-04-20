package article_entity

import (
	"time"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

const MAX_SUMMARY_LENGTH = 100

// should change those urls into url.URL type
type Article struct {
	ID        uint        `gorm:"type:autoIncrement" json:"id"`
	AuthorID  uint        `json:"authorId"`
	CreatedAt time.Time   `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime" json:"updatedAt"`
	Summary   string      `json:"summary"`
	Thumbnail string      `json:"thumbnail"`
	Images    []string    `json:"-"`
	Body      ArticleBody `gorm:"constraint:OnDelete:CASCADE" json:"-"`
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

func (a *Article) SetSummary(body []byte) *Article {
	plainText := util.SanitizeExceptPlainText(body)
	minLength := util.GetMin(len(plainText), MAX_SUMMARY_LENGTH)
	a.Summary = string(plainText[:minLength])
	return a
}

func (a *Article) SetImages(body []byte) *Article {
	a.Images = util.ExtractImageURL(body)
	a.setThumbnail()
	return a
}

func (a *Article) setThumbnail() {
	if len(a.Images) == 0 {
		a.Thumbnail = ""
	} else {
		a.Thumbnail = a.Images[0]
	}
}
