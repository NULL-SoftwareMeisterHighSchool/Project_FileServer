package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

func CreateArticle(authorID uint, title string, articleType article_entity.ArticleType) error {
	article := article_entity.New().
		SetAuthorID(authorID).
		SetTitle(title).
		SetArticleType(articleType)

	tx := database.Articles.Create(&article)
	return tx.Error
}

func DeleteByID(id uint) {
	database.Articles.Where("id = ?", id).Delete(article_entity.New())
}

func DeleteByAuthorID(authorID uint) {
	database.Articles.Where("authorID = ?", authorID).Delete(article_entity.New())
}

func GetArticleWithBody(id uint) *article_entity.Article {
	var article *article_entity.Article
	database.Articles.Where("id = ?", id).Association("body").Find(article)
	return article
}
