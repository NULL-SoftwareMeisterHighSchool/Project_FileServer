package articles

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
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

	articles := db.GetArticleInfoByIDs(ids)
	return c.Status(http.StatusOK).JSON(articles)
}

func GetArticle(c *fiber.Ctx) error {
	id, authorID := getArticleIDAndAuthorID(c)
	storage := storages.Get()

	article, size := storage.GetArticle(authorID, id)
	if article == nil {
		return errors.ErrArticleNotFound
	}

	return c.Status(http.StatusOK).SendStream(article, int(size))
}

func CreateArticle(c *fiber.Ctx) error {
	id, authorID := getArticleIDAndAuthorID(c)
	storage := storages.Get()

	if storage.ArticleExists(authorID, id) {
		return errors.ErrConflict
	}

	body := util.SanitizeXSS(c.Body())
	article := db.CreateArticle(id, authorID, body)
	db.Save(article)

	if err := storage.WriteArticle(authorID, id, body); err != nil {
		return err
	}
	return c.SendStatus(http.StatusCreated)
}

func UpdateArticle(c *fiber.Ctx) error {
	id, authorID := getArticleIDAndAuthorID(c)
	storage := storages.Get()
	body := util.SanitizeXSS(c.Body())

	if !storage.ArticleExists(authorID, id) {
		return errors.ErrArticleNotFound
	}

	newArticle := db.CreateArticle(id, authorID, body)
	diff := util.GetDifferenceBetweenStrArr(
		db.GetImageURLsByID(id), newArticle.Images,
	)
	imagesToDelete := filterDeletableImageURLByEndpoint(diff, config.IMAGE_HOST_ENDPOINT)
	storage.DeleteImages(authorID, storage.GetSuffixesFromURLs(imagesToDelete))
	db.Save(newArticle)

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
		db.GetImageURLsByID(id), config.IMAGE_HOST_ENDPOINT,
	)
	storage.DeleteImages(authorID, storage.GetSuffixesFromURLs(imagesToDelete))
	db.DeleteByID(id)

	return c.SendStatus(http.StatusNoContent)
}
