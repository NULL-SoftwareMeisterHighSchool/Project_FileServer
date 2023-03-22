package articles

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetArticle(c *fiber.Ctx) error {
	id := c.Locals("id").(int)
	author := c.Locals("author").(string)

	return c.SendFile(getArticlePath(author, id))
}

func CreateArticle(c *fiber.Ctx) error {
	id := c.Locals("id").(int)
	author := c.Locals("author").(string)
		
	return c.SendStatus(http.StatusCreated)
}

func UpdateArticle(c *fiber.Ctx) error {
	// id := c.Locals("id").(int)
	// author := c.Locals("author").(string)
	return c.SendStatus(http.StatusAccepted)
}

func DeleteArticle(c *fiber.Ctx) error {
	// id := c.Locals("id").(int)
	// author := c.Locals("author").(string)
	return c.SendStatus(http.StatusNoContent)
}
