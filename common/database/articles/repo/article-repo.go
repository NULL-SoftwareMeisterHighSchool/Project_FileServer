package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

func Save(article *article_entity.Article) {
	database.DB.Save(article)
}

// CreateArticle creates article from body. but doesn't save it
func CreateArticle(articleID, authorID uint, body []byte) *article_entity.Article {
	return article_entity.New().
		SetAuthorID(authorID).
		SetID(articleID).
		SetSummary(body).
		SetImages(body)
}

func DeleteByID(id uint) {
	database.DB.Where("id = ?", id).Delete(article_entity.New())
}

func DeleteByAuthorID(authorID uint) {
	database.DB.Where("userId = ?", authorID).Delete(article_entity.New())
}

func GetImageURLsByID(id uint) []string {
	urls := []string{}
	database.DB.Model(article_entity.New()).Where("id = ?", id).Select("images").Take(urls)
	return urls
}

func GetArticleInfoByIDs(ids []uint) []*article_entity.Article {
	var articles []*article_entity.Article
	database.DB.
		Where("id IN ?", ids).
		Omit("images").
		Find(&articles).
		Order("id")

	return articles
}
