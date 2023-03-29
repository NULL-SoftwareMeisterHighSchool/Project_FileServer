package articles

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func getArticlePath(author string, id uint) string {
	return fmt.Sprintf("./contents/%s/articles/%d.md", author, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getIdAndAuthor(c *fiber.Ctx) (id uint, author string) {
	id = c.Locals("id").(uint)
	author = c.Locals("author").(string)
	return
}
