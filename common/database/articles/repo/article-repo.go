package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

func Save(article *article_entity.Article) {
	database.Articles.Save(article)
}

// CreateArticle creates article from body. but doesn't save it
func CreateArticle(articleID, authorID uint, body []byte, images []string) *article_entity.Article {
	return article_entity.New().
		SetAuthorID(authorID).
		SetID(articleID).
		SetBody(body).
		SetSummary(body).
		SetImagesAndThumbnail(images)
}

func DeleteByID(id uint) {
	database.Articles.Where("id = ?", id).Delete(article_entity.New())
}

func DeleteByAuthorID(authorID uint) {
	database.Articles.Where("userId = ?", authorID).Delete(article_entity.New())
}

func GetImageURLsByID(id uint) []string {
	urls := []string{}
	database.Articles.Where("id = ?", id).Select("images").Take(urls)
	return urls
}

func GetArticleInfoByIDs(ids []uint) []*article_entity.Article {
	var articles []*article_entity.Article
	database.Articles.
		Where("id IN ?", ids).
		Omit("images").
		Find(&articles).
		Order("id")

	return articles
}

func GetArticleWithBody(id uint) *article_entity.Article {
	var article *article_entity.Article
	database.Articles.Where("id = ?", id).Association("body").Find(article)
	return article
}
