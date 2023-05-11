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

func (fs) DeleteImages(urls []string) {
	for _, url := range urls {
		os.Remove(getImagePath(getSuffixFromURL(url)))
	}
}

func (fs) UploadImage(name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error) {
	fullName := strings.Join([]string{name, extension}, ".")
	path := getImagePath(fullName)

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
	return getImageURL(fullName), nil
}
