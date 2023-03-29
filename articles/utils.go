package articles

import (
	"github.com/gofiber/fiber/v2"
)

func getIdAndAuthor(c *fiber.Ctx) (id uint, author string) {
	id = c.Locals("id").(uint)
	author = c.Locals("author").(string)
	return
}
