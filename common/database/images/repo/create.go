package image_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	image_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/images/entity"
)

func CreateImage(articleID uint, url string) error {
	image := image_entity.Image{URL: url}

	database.Articles.Where("id = ?", articleID).
		Association("images").
		Append(&image)

	return nil
}
