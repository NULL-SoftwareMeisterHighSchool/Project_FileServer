package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	"gorm.io/gorm"
)

func IncreaseViewCount(articleID uint) {
	// might need to seperate the views into table. someday
	database.Articles.
		Where("id = ?", articleID).
		UpdateColumn("views", gorm.Expr("views  + ?", 1))
}

func UpdateArticleBodyAndImages(articleID uint, body []byte, images []string) {

	article := article_entity.New().
		SetID(articleID).
		SetBody(body).
		SetSummary(body).
		SetImagesAndThumbnail(images)

	database.Articles.
		Where("id = ?", articleID).
		Omit("id").
		Updates(article)
}

func UpdateTitleByID(articleID uint, title string) {
	database.Articles.
		Where("id = ?", articleID).
		Update("title", title)
}

func UpdateIsPrivateByID(articleID uint, isPrivate bool) {
	database.Articles.
		Where("id = ?", articleID).
		Update("isPrivate", isPrivate)
}

func ToggleLike(articleID, userID uint) {

	var exists bool
	it := map[string]interface{}{
		"article_id": articleID,
		"user_id":    userID,
	}

	database.ArticleLikes.
		Where(it).
		Select("COUNT(*) > 0").Find(&exists)

	if exists {
		database.ArticleLikes.Delete(it)
	} else {
		database.ArticleLikes.Create(it)
	}
}
