package filesystem

import (
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

type fs struct{}

func Get() *fs {
	return &fs{}
}

func (fs) DeleteImages(authorID uint, suffixes []string) {
	for _, suffix := range suffixes {
		os.Remove(getImagePath(authorID, suffix))
	}
}

func (fs) GetSuffixesFromURLs(urls []string) []string {
	suffixes := []string{}
	for _, url := range urls {
		splitted := strings.Split(url, "/")
		suffix := splitted[len(splitted)-1]
		suffixes = append(suffixes, suffix)
	}
	return suffixes
}

func (fs) UploadImage(authorID uint, name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error) {
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
	return getImageURL(authorID, fullName), nil
}
