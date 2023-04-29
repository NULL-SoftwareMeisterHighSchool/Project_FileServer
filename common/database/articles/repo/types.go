package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

type ArticleOrder uint8

const (
	TIME ArticleOrder = iota
	POPULARITY
	RELEVANCE
)

type ArticleTypeOption uint8

const (
	ALL ArticleTypeOption = iota
	GENERAL
	TECH
)

func DeleteByID(id uint) {
	database.Articles.Delete(article_entity.New().SetID(id))
}

func DeleteByAuthorID(authorID uint) {
	database.Articles.Delete(article_entity.New().SetAuthorID(authorID))
}

func GetArticleWithBody(id uint) *article_entity.Article {
	var article *article_entity.Article
	database.Articles.Where("id = ?", id).Association("body").Find(article)
	return article
}
