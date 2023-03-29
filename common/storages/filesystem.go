package storages

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

type filesystem struct{}

func (filesystem) ArticleExists(author string, id uint) bool {
	return articleExistsByPath(getArticlePath(author, id))
}

func (filesystem) WriteArticle(author string, id uint, body []byte) *fiber.Error {
	articlePath := getArticlePath(author, id)
	if err := os.WriteFile(articlePath, body, 0666); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (filesystem) GetArticle(author string, id uint) []byte {
	articlePath := getArticlePath(author, id)
	bytes, err := os.ReadFile(articlePath)
	if err != nil {
		return nil
	}
	return bytes
}

func (filesystem) DeleteArticle(author string, id uint) *fiber.Error {
	articlePath := getArticlePath(author, id)
	if err := os.Remove(articlePath); err != nil {
		return errors.NotFoundError
	}
	return nil
}

func (filesystem) DeleteImages(urls []string) {

}

func getArticlePath(author string, id uint) string {
	return fmt.Sprintf("./contents/%s/articles/%d.md", author, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
