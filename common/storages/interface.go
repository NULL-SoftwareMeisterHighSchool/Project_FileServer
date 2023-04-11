package storages

import (
	"io"
	"mime/multipart"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages/filesystem"
	"github.com/gofiber/fiber/v2"
)

type Storage interface {
	ArticleExists(authorID uint, id uint) bool
	WriteArticle(authorID uint, id uint, body []byte) *fiber.Error
	GetArticle(authorID uint, id uint) (io.Reader, int64)
	DeleteArticle(authorID uint, id uint) *fiber.Error
	DeleteImages(authorID uint, suffixes []string)
	GetSuffixesFromURLs(urls []string) []string
	UploadImage(authorID uint, name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error)
}

var storage Storage = filesystem.Get()

func Get() Storage {
	return storage
}
