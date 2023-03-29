package articles

import (
	"fmt"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/gofiber/fiber/v2"
)

func getIdAndAuthor(c *fiber.Ctx) (id uint, author string) {
	id = c.Locals("id").(uint)
	author = c.Locals("author").(string)
	return
}

func filterDeletableImageURLs(urls []string) []string {
	shouldDelete := []string{}
	for _, url := range urls {
		if strings.HasPrefix(url, fmt.Sprintf("https://%s/images/", config.HOSTNAME)) {
			shouldDelete = append(shouldDelete, url)
		}
	}
	return shouldDelete
}
