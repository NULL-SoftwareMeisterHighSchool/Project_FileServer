package article_repo

import (
	"time"

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

type ListArticleElemWithLikes struct {
	article_entity.Article
	Likes uint
}

func CreateArticle(authorID uint, title string, articleType article_entity.ArticleType) error {
	article := article_entity.New().
		SetAuthorID(authorID).
		SetTitle(title).
		SetArticleType(articleType).
		SetIsPrivate(true)

	tx := database.Articles.Create(&article)
	return tx.Error
}

func ListArticles(offset uint,
	amount uint,
	order ArticleOrder,
	articleType ArticleTypeOption,
	authorID uint,
	isPrivate *bool,
	start time.Time,
	end time.Time,
	query string) []*ListArticleElemWithLikes {

	var articles []*ListArticleElemWithLikes

	tx := database.Articles
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
		tx = tx.Where("MATCH(title) AGAINST (?)", query).Or("id IN (?)",
			database.ArticleBodies.
				Where("MATCH(text) AGAINST (?)", query).Select("article_id"),
		)
	}

	tx = tx.Select("articles.*, (?) as likes",
		database.ArticleLikes.Where("article_id = articles.id").
			Select("COUNT(*)"),
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
