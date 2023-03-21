package articles

import (
	"github.com/gofiber/fiber/v2"
)

func GetArticle(c *fiber.Ctx) error {
	id := c.Locals("id").(int)
	author := c.Locals("author").(string)

	return c.SendFile(getArticlePath(author, id))
}

func CreateArticle(c *fiber.Ctx) error {
	// id := c.Locals("id").(int)
	// author := c.Locals("author").(string)
	return nil
}

func UpdateArticle(c *fiber.Ctx) error {
	// id := c.Locals("id").(int)
	// author := c.Locals("author").(string)
	return nil
}

func DeleteArticle(c *fiber.Ctx) error {
	// id := c.Locals("id").(int)
	// author := c.Locals("author").(string)
	return nil
}
