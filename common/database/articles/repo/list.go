package article_repo

import (
	"time"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	"gorm.io/gorm"
)

type ListArticleElemWithLikes struct {
	article_entity.Article
	Thumbnail string
	Likes     uint
	Comments  uint
	IsAuthor  bool
}

func ListArticles(
	offset uint,
	amount uint,
	order ArticleOrder,
	articleType ArticleTypeOption,
	authorID uint,
	userID uint,
	isPrivate *bool,
	start time.Time,
	end time.Time,
	query string,
) ([]*ListArticleElemWithLikes, uint) {

	var articles []*ListArticleElemWithLikes
	var recordCnt int64

	tx := database.Articles()

	tx = tx.Where("created_at BETWEEN ? AND ?", start, end)

	tx = tx.Where("type = ?", articleType)

	if authorID != 0 {
		tx = tx.Where("author_id = ?", authorID)
	}

	if isPrivate != nil {
		tx = tx.Where("is_private = ?", *isPrivate)
	}

	if query != "" {
		tx = tx.
			Where(tx.Scopes(
				func(d *gorm.DB) *gorm.DB {
					return tx.Where("(MATCH(title) AGAINST (?)", query)
				},
				func(d *gorm.DB) *gorm.DB {
					return tx.Or("id IN (?))",
						database.ArticleBodies().
							Where("MATCH(text) AGAINST (?)", query).
							Select("article_id"),
					)
				}))
	}

	// get count
	tx.Count(&recordCnt)

	// select
	tx = tx.Select("articles.*, (?) AS likes,  (?) AS comments, articles.author_id = ? AS is_author",
		LikesForArticleQuery().
			Select("COUNT(*)"),
		// comments
		database.Comments().
			Where("article_id = articles.id").
			Select("COUNT(*)"),
		// userID
		userID,
	)

	switch order {
	case TIME:
		tx = tx.Order("created_at DESC")
	case VIEWS:
		tx = tx.Order("views DESC")
	case LIKES:
		tx = tx.Order("likes DESC")
	}

	tx.Offset(int(offset)).Limit(int(amount)).Find(&articles)

	return articles, uint(recordCnt)
}
