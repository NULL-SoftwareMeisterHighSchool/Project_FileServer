package image_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	image_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/images/entity"
)

func CreateImage(url string) error {
	image := image_entity.Image{
		URL: url,
	}

	database.Images().Create(&image)
	return nil
}
