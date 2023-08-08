package article_entity

import (
	"time"

	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

const MAX_SUMMARY_LENGTH = 400

type ArticleType uint8

const (
	TYPE_GENERAL ArticleType = iota
	TYPE_TECH
)

type Article struct {
	ID        uint              `gorm:"autoIncrement"`
	AuthorID  uint              `gorm:"not null"`
	Author    *user_entity.User `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time         `gorm:"autoCreateTime"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime"`
	Title     string            `gorm:"type:varchar(2048);index:,class:FULLTEXT,option:WITH PARSER ngram"`
	Summary   string            `gorm:"type:varchar(400)"`
	Body      ArticleBody       `gorm:"constraint:OnDelete:CASCADE"`
	Type      ArticleType       `gorm:"type:tinyint"`
	IsPrivate bool              `gorm:"type:tinyint"`
	Views     uint64
	Likes     []*user_entity.User       `gorm:"many2many:user_likes;constraint:OnDelete:CASCADE"`
	Comments  []*comment_entity.Comment `gorm:"constraint:OnDelete:CASCADE;"`
}

type ArticleBody struct {
	ArticleID uint
	Text      string `gorm:"type:longtext;index:,class:FULLTEXT,option:WITH PARSER ngram"`
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

func (a *Article) SetArticleType(articleType ArticleType) *Article {
	a.Type = articleType
	return a
}

func (a *Article) SetIsPrivate(isPrivate bool) *Article {
	a.IsPrivate = isPrivate
	return a
}
