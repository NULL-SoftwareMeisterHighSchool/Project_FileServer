package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
)

func DeleteByID(id uint) {
	database.Articles.Delete(article_entity.New().SetID(id))
}

func DeleteByAuthorID(authorID uint) {
	database.Articles.Delete(article_entity.New().SetAuthorID(authorID))
}
