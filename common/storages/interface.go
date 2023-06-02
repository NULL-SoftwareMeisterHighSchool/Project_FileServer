package storages

import (
	"mime/multipart"
	"sync"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages/s3"
	"github.com/gofiber/fiber/v2"
)

type Storage interface {
	DeleteImages(urls []string)
	UploadImage(name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error)
}

var storage Storage
var once sync.Once

func Get() Storage {
	once.Do(func() {
		storage = s3.Get()
	})
	return storage
}
