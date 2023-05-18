package errors

import (
	"fmt"
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/gofiber/fiber/v2"
)

var ErrArticleNotFound = fiber.NewError(http.StatusNotFound, "article does not exist")
var ErrConflict = fiber.NewError(http.StatusConflict, "article already exists")
var ErrAuthFailed = fiber.NewError(http.StatusUnauthorized, "invalid auth method or token")
var ErrInvalidID = fiber.NewError(http.StatusBadRequest, "id should be valid")

var ErrImageNotFound = fiber.NewError(http.StatusNotFound, "image does not exist")
var ErrInvalidImageExtension = fiber.NewError(
	http.StatusUnsupportedMediaType,
	fmt.Sprintf("allowed file extensions are: %v", config.IMAGE_EXTENSIONS),
)
var ErrImageNotReceived = fiber.NewError(http.StatusBadRequest, "image not received")
var ErrInvalidArticleID = fiber.NewError(http.StatusBadRequest, "articleID from querystring isn't valid")

func ErrNoPermission(err error) *fiber.Error {
	return fiber.NewError(http.StatusForbidden, "author and requested user did not match: "+err.Error())
}

func CreateUnkownErr(e error) *fiber.Error {
	return fiber.NewError(http.StatusInternalServerError, e.Error())
}
