package errors

import (
	"fmt"
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/gofiber/fiber/v2"
)

var ArticleNotFoundError = fiber.NewError(http.StatusNotFound, "article does not exist")
var ConflictError = fiber.NewError(http.StatusConflict, "article already exists")
var AuthFailedError = fiber.NewError(http.StatusUnauthorized, "invalid auth method or token")
var NoPermissionError = fiber.NewError(http.StatusForbidden, "author and requested user did not match")
var InvalidIDError = fiber.NewError(http.StatusBadRequest, "id should be valid")

var ImageNotFoundError = fiber.NewError(http.StatusNotFound, "image does not exist")
var InvalidImageExtensionError = fiber.NewError(
	http.StatusUnsupportedMediaType,
	fmt.Sprintf("allowed file extensions are: %v", config.IMAGE_EXTENSIONS),
)
var ImageNotReceived = fiber.NewError(http.StatusBadRequest, "image not received")

func CreateUnkownErr(e error) *fiber.Error {
	return fiber.NewError(http.StatusInternalServerError, e.Error())
}
