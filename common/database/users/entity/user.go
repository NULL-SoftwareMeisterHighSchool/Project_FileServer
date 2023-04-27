package user_entity

import article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"

type User struct {
	ID               uint
	ArticlesUploaded []*article_entity.Article `gorm:"foreignKey:AuthorID"`
	ArticlesLiked    []*article_entity.Article `gorm:"many2many:user_likes;"`
}
