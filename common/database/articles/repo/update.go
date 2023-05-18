package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	"gorm.io/gorm"
)

func IncreaseViewCount(articleID uint) {
	// might need to seperate the views into table. someday
	database.Articles().
		Where("id = ?", articleID).
		UpdateColumn("views", gorm.Expr("views  + ?", 1))
}

func UpdateArticleBody(articleID uint, body []byte) error {

	article := article_entity.New().
		SetID(articleID).
		SetBody(body).
		SetSummary(body)

	tx := database.Articles().
		Where("id = ?", articleID).
		Omit("id").
		Updates(article)
	return tx.Error
}

func UpdateTitleByID(articleID uint, title string) error {
	tx := database.Articles().
		Where("id = ?", articleID).
		Update("title", title)
	return tx.Error
}

func UpdateIsPrivateByID(articleID uint, isPrivate bool) error {
	tx := database.Articles().
		Where("id = ?", articleID).
		Update("is_private", isPrivate)
	return tx.Error
}

func ToggleLike(articleID, userID uint) error {

	var exists bool
	it := map[string]interface{}{
		"article_id": articleID,
		"user_id":    userID,
	}

	database.ArticleLikes().
		Where(it).
		Select("COUNT(*) > 0").Find(&exists)

	if exists {
		return database.ArticleLikes().Delete(it).Error
	} else {
		return database.ArticleLikes().Create(it).Error
	}
}
