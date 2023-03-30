package storages

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

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

func (filesystem) GetArticle(author string, id uint) (io.Reader, int64) {
	articlePath := getArticlePath(author, id)
	file, err := os.Open(articlePath)
	stat, err := file.Stat()
	if err != nil {
		return nil, -1
	}
	return file, stat.Size()
}

func (filesystem) DeleteArticle(author string, id uint) *fiber.Error {
	articlePath := getArticlePath(author, id)
	if err := os.Remove(articlePath); err != nil {
		return errors.ArticleNotFoundError
	}
	return nil
}

func (filesystem) DeleteImages(author string, suffixes []string) {
	for _, suffix := range suffixes {
		os.Remove(getImagePath(author, suffix))
	}
}

func (filesystem) GetSuffixesFromURLs(urls []string) []string {
	suffixes := []string{}
	for _, url := range urls {
		suffix := fmt.Sprint(strings.Split(url, "images/")[1])
		suffixes = append(suffixes, suffix)
	}
	return suffixes
}

func (filesystem) UploadImage(file *multipart.FileHeader) *fiber.Error {
	
}

func getImagePath(author, suffix string) string {
	return fmt.Sprintf("./contents/images/%s/%s", author, suffix)
}

func getArticlePath(author string, id uint) string {
	return fmt.Sprintf("./contents/articles/%s/%d.md", author, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
