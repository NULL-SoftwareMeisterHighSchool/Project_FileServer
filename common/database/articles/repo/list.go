package article_repo

import (
	"time"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

type ListArticleElemWithLikes struct {
	article_entity.Article
	Thumbnail string
	Likes     uint
}

func ListArticles(
	offset uint,
	amount uint,
	order ArticleOrder,
	articleType ArticleTypeOption,
	authorID uint,
	isPrivate *bool,
	start time.Time,
	end time.Time,
	query string,
) []*ListArticleElemWithLikes {

	var articles []*ListArticleElemWithLikes
	tx := database.Articles()

	tx = tx.Where("created_at BETWEEN ? AND ?", start, end)

	if articleType != ALL {
		tx = tx.Where("type = ?", articleType-1)
	}

	if authorID != 0 {
		tx = tx.Where("author_id = ?", authorID)
	}

	if isPrivate != nil {
		tx = tx.Where("is_private = ?", *isPrivate)
	}

	if query != "" {
		tx = tx.
			Where("MATCH(title) AGAINST (?)", query).
			Or("id IN (?)",
				database.ArticleBodies().
					Where("MATCH(text) AGAINST (?)", query).
					Select("article_id"),
			)
	}

	tx = tx.Select("articles.*, (?) AS likes, (?) AS thumbnail",
		LikesForArticleQuery().
			Select("COUNT(*)"),
		database.Images().
			Where("article_id = articles.id").
			Select("url").
			Limit(1),
	)

	// tx = tx.Joins("JOIN (?) AS likes_tbl ON articles.id = likes_tbl.article_id",
	// 	database.ArticleLikes.Group("article_id").
	// )

	switch order {
	case TIME:
		tx = tx.Order("created_at DESC")
	case POPULARITY:
		tx = tx.Order("views + (likes * 3) DESC")
	}

	tx.Offset(int(offset)).Limit(int(amount)).Find(&articles)

	return articles
}
