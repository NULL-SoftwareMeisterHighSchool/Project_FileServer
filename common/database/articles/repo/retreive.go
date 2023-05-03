package article_repo

import (
	"errors"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	"gorm.io/gorm"
)

type ArticlePermissionInfo struct {
	AuthorID  uint
	IsPrivate bool
}

type ArticleWithBodyAndLikes struct {
	article_entity.Article
	Likes   uint
	IsLiked bool
}

func GetArticlePermissionInfoByID(id uint) (*ArticlePermissionInfo, error) {
	var info *ArticlePermissionInfo
	err := database.Articles.
		Where("id = ?", id).
		First(info).
		Error
	return info, err
}

func GetArticleWithBody(id, userID uint) (*ArticleWithBodyAndLikes, error) {
	var article ArticleWithBodyAndLikes

	tx := database.Articles.Where("id = ?", id)

	tx = tx.Select("articles.*, (?) AS likes, (?) AS is_liked",
		likesForArticleQuery.
			Select("COUNT(*)"),
		likesForArticleQuery.
			Where("user_id = ? AND user_id != 0", userID).
			Select("COUNT(*) > 0"),
	)

	tx = tx.Joins("ArticleBody").Omit("summary", "thumbnail", "images").First(&article)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	return &article, nil
}
