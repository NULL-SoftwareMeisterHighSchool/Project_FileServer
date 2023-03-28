package articles

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func getArticlePath(author string, id int) string {
	return fmt.Sprintf("./contents/%s/articles/%d.md", author, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getIdAndAuthor(c *fiber.Ctx) (id int, author string) {
	id = c.Locals("id").(int)
	author = c.Locals("author").(string)
	return
}
