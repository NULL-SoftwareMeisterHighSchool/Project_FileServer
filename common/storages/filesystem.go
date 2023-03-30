package storages

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
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
		splitted := strings.Split(url, "/")
		suffix := splitted[len(splitted)-1]
		suffixes = append(suffixes, suffix)
	}
	return suffixes
}

func (filesystem) UploadImage(author, name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error) {
	fullName := strings.Join([]string{name, extension}, ".")
	path := getImagePath(author, fullName)

	receivedFile, _ := fileHeader.Open()
	defer receivedFile.Close()
	createdFile, err := os.Create(path)
	if err != nil {
		return "", errors.CreateUnkownErr(err)
	}
	defer createdFile.Close()

	if _, err = io.Copy(createdFile, receivedFile); err != nil {
		return "", errors.CreateUnkownErr(err)
	}
	return getFSArticleURL(author, fullName), nil
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

func getFSArticleURL(author, fullName string) string {
	return fmt.Sprintf("%s/%s/%s", config.IMAGE_HOST_ENDPOINT, author, fullName)
}
