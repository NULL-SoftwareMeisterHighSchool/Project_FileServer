package storages

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type filesystem struct{}

func (_ filesystem) ArticleExists(author string, id uint) bool {
	return articleExistsByPath(getArticlePath(author, id))
}
func (_ filesystem) WriteArticle(author string, id uint, body []byte) *fiber.Error {

}
func (_ filesystem) GetArticle(author string, id uint) []byte {

}
func (_ filesystem) DeleteArticle(author string, id uint) *fiber.Error {

}

func getArticlePath(author string, id uint) string {
	return fmt.Sprintf("./contents/%s/articles/%d.md", author, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
