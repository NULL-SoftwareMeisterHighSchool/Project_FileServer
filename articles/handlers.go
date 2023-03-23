package articles

import (
	"net/http"
	"os"

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

	articlePath := getArticlePath(author, id)
	if articleExistsByPath(articlePath) {
		return fiber.NewError(http.StatusConflict, "article already exists")
	}

	// TODO : should create article entity. which contains image info
	if err := os.WriteFile(articlePath, c.Body(), 0666); err != nil {
		return err
	}
	return c.SendStatus(http.StatusCreated)
}

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Locals("id").(int)
	author := c.Locals("author").(string)

	articlePath := getArticlePath(author, id)
	if err := os.WriteFile(articlePath, c.Body(), 0666); err != nil {
		return err
	}

	// TODO : should update article entity, and images
	return c.SendStatus(http.StatusAccepted)
}

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Locals("id").(int)
	author := c.Locals("author").(string)

	articlePath := getArticlePath(author, id)
	if err := os.Remove(articlePath); err != nil {
		return fiber.NewError(http.StatusNotFound, "article does not exist")
	}
	// TODO : should delete article entity, and images
	
	return c.SendStatus(http.StatusNoContent)
}
