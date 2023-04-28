package articles

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
)

func GetArticleInfo(c *fiber.Ctx) error {
	ids := getIdsFromQuery(c.Query("ids"))
	if len(ids) == 0 {
		return errors.ErrInvalidID
	}

	articles := db.DB().GetArticleInfoByIDs(ids)
	return c.Status(http.StatusOK).JSON(articles)
}

func GetArticle(c *fiber.Ctx) error {
	id, authorID := getArticleIDAndAuthorID(c)
	article_repo.
		article, size := storage.GetArticle(authorID, id)
	if article == nil {
		return errors.ErrArticleNotFound
	}

	return c.Status(http.StatusOK).SendStream(article, int(size))
}

func UpdateArticle(c *fiber.Ctx) error {
	id, authorID := getArticleIDAndAuthorID(c)
	storage := storages.Get()
	body := util.SanitizeXSS(c.Body())

	if !storage.ArticleExists(authorID, id) {
		return errors.ErrArticleNotFound
	}

	newArticle := db.DB().CreateArticle(id, authorID, body)
	diff := util.GetDifferenceBetweenStrArr(
		db.DB().GetImageURLsByID(id), newArticle.Images,
	)
	imagesToDelete := filterDeletableImageURLByEndpoint(diff, config.IMAGE_HOST_ENDPOINT)
	go storage.DeleteImages(authorID, storage.GetSuffixesFromURLs(imagesToDelete))
	go db.DB().Save(newArticle)

	if err := storage.WriteArticle(authorID, id, body); err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}

func DeleteArticle(c *fiber.Ctx) error {
	id, authorID := getArticleIDAndAuthorID(c)
	storage := storages.Get()

	if err := storage.DeleteArticle(authorID, id); err != nil {
		return err
	}

	imagesToDelete := filterDeletableImageURLByEndpoint(
		db.DB().GetImageURLsByID(id), config.IMAGE_HOST_ENDPOINT,
	)
	go storage.DeleteImages(authorID, storage.GetSuffixesFromURLs(imagesToDelete))
	go db.DB().DeleteByID(id)

	return c.SendStatus(http.StatusNoContent)
}
