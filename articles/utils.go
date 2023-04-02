package articles

import (
	"strconv"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/gofiber/fiber/v2"
)

func getArticleIDAndAuthorID(c *fiber.Ctx) (id uint, authorID uint) {
	id = c.Locals("id").(uint)
	authorID = c.Locals("authorID").(uint)
	return
}

func filterDeletableImageURLs(urls []string) []string {
	shouldDelete := []string{}
	for _, url := range urls {
		if strings.HasPrefix(url, config.IMAGE_HOST_ENDPOINT) {
			shouldDelete = append(shouldDelete, url)
		}
	}
	return shouldDelete
}

func getIdsFromQuery(idArrStr string) []uint {
	var ids []uint
	idArr := strings.Split(idArrStr, ",")
	for _, idStr := range idArr {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return make([]uint, 0)
		}
		ids = append(ids, uint(id))
	}
	return ids
}
