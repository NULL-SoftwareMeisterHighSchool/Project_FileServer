package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func getParam(to *string, name string, c *fiber.Ctx) *fiber.Error {
	*to = c.Params("id")
	if *to == "" {
		return fiber.NewError(http.StatusNotFound, "id should be valid")
	}
	return nil
}
