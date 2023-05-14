package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	"gorm.io/gorm"
)

func LikesForArticleQuery() *gorm.DB {
	return database.ArticleLikes.Where("article_id = articles.id")
}
