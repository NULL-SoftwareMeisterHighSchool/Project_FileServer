package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

func CreateArticle(authorID uint, title string, articleType article_entity.ArticleType) error {
	article := article_entity.New().
		SetAuthorID(authorID).
		SetTitle(title).
		SetArticleType(articleType).
		SetIsPrivate(true)

	tx := database.Articles.Create(&article)
	return tx.Error
}
