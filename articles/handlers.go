package articles

import (
	"github.com/gofiber/fiber/v2"
)

func GetArticle(c *fiber.Ctx) error {
	id := c.Locals("id").(string)
	username := c.Locals("username").(string)

	return c.SendFile(getArticlePath(username, id))
}

func CreateArticle(c *fiber.Ctx) error {
	id := c.Locals("id")
}

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Locals("id")
}

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Locals("id")
}
