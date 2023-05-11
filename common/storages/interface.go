package storages

import (
	"mime/multipart"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages/filesystem"
	"github.com/gofiber/fiber/v2"
)

type Storage interface {
	DeleteImages(names []string)
	UploadImage(name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error)
}

var storage Storage = filesystem.Get()

func Get() Storage {
	return storage
}
