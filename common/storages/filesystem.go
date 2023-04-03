package storages

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

type filesystem struct{}

func (filesystem) ArticleExists(authorID uint, id uint) bool {
	return articleExistsByPath(getArticlePath(authorID, id))
}

func (filesystem) WriteArticle(authorID uint, id uint, body []byte) *fiber.Error {
	articlePath := getArticlePath(authorID, id)
	if err := os.WriteFile(articlePath, body, 0666); err != nil {
		return errors.CreateUnkownErr(err)
	}
	return nil
}

func (filesystem) GetArticle(authorID uint, id uint) (io.Reader, int64) {
	articlePath := getArticlePath(authorID, id)
	file, _ := os.Open(articlePath)
	stat, err := file.Stat()
	if err != nil {
		return nil, -1
	}
	return file, stat.Size()
}

func (filesystem) DeleteArticle(authorID uint, id uint) *fiber.Error {
	articlePath := getArticlePath(authorID, id)
	if err := os.Remove(articlePath); err != nil {
		return errors.ErrArticleNotFound
	}
	return nil
}

func (filesystem) DeleteImages(authorID uint, suffixes []string) {
	for _, suffix := range suffixes {
		os.Remove(getImagePath(authorID, suffix))
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

func (filesystem) UploadImage(authorID uint, name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error) {
	fullName := strings.Join([]string{name, extension}, ".")
	path := getImagePath(authorID, fullName)

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
	return getFSArticleURL(authorID, fullName), nil
}

func getImagePath(authorID uint, suffix string) string {
	return fmt.Sprintf("./contents/images/%d/%s", authorID, suffix)
}

func getArticlePath(authorID uint, id uint) string {
	return fmt.Sprintf("./contents/articles/%d/%d.md", authorID, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getFSArticleURL(authorID uint, fullName string) string {
	return fmt.Sprintf("%s/%d/%s", config.IMAGE_HOST_ENDPOINT, authorID, fullName)
}
