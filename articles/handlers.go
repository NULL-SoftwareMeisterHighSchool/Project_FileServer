package articles

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
)

func GetArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	storage := storages.Get()

	article, size := storage.GetArticle(author, id)
	if article == nil {
		return errors.ArticleNotFoundError
	}

	return c.Status(http.StatusOK).SendStream(article, int(size))
}

func CreateArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	storage := storages.Get()

	if storage.ArticleExists(author, id) {
		return errors.ConflictError
	}

	body := util.SanitizeXSS(c.Body())
	article := db.CreateArticle(id, body)
	db.Save(article)

	if err := storage.WriteArticle(author, id, body); err != nil {
		return err
	}
	return c.SendStatus(http.StatusCreated)
}

func UpdateArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	storage := storages.Get()
	body := util.SanitizeXSS(c.Body())

	if !storage.ArticleExists(author, id) {
		return errors.ArticleNotFoundError
	}

	newArticle := db.CreateArticle(id, body)
	diff := util.GetDifferenceBetweenStrArr(
		db.GetImageURLsByID(id), newArticle.Images,
	)
	imagesToDelete := filterDeletableImageURLs(diff)
	storage.DeleteImages(author, storage.GetSuffixesFromURLs(imagesToDelete))
	db.Save(newArticle)

	if err := storage.WriteArticle(author, id, body); err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}

func DeleteArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	storage := storages.Get()

	if err := storage.DeleteArticle(author, id); err != nil {
		return err
	}

	imagesToDelete := filterDeletableImageURLs(db.GetImageURLsByID(id))
	storage.DeleteImages(author, storage.GetSuffixesFromURLs(imagesToDelete))
	db.DeleteByID(id)

	return c.SendStatus(http.StatusNoContent)
}
