package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

func DeleteByID(id uint) error {
	tx := database.Articles().Delete(article_entity.New().SetID(id))
	return tx.Error
}

func DeleteByAuthorID(authorID uint) error {
	tx := database.Articles().Delete(article_entity.New().SetAuthorID(authorID))
	return tx.Error
}
