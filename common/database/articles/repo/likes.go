package article_repo

import "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"

var likesForArticleQuery = database.ArticleLikes.Where("article_id = articles.id")
