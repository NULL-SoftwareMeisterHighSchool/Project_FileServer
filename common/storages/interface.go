package storages

import "github.com/gofiber/fiber/v2"

type Storage interface {
	ArticleExists(author string, id uint) bool
	WriteArticle(author string, id uint, body []byte) *fiber.Error
	GetArticle(author string, id uint) []byte
	DeleteArticle(author string, id uint) *fiber.Error
	DeleteImages(urls []string)
}

var storage = &filesystem{}

func Get() Storage {
	return storage
}