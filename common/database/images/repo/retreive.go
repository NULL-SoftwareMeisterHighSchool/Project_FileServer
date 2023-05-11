package image_repo

import "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"

func GetURLsByArticleID(articleID uint) []string {
	var urls []string

	database.Images.
		Where("article_id = ?", articleID).
		Select("url").
		Find(&urls)

	return urls
}
