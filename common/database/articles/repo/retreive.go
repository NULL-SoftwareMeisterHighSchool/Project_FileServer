package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
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
	var info = &ArticlePermissionInfo{}
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
		LikesForArticleQuery().
			Select("COUNT(*)"),
		LikesForArticleQuery().
			Where("user_id = ? AND user_id != 0", userID).
			Select("COUNT(*) > 0"),
	)

	tx = tx.Joins("ArticleBody").Omit("summary", "images").First(&article)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &article, nil
}
