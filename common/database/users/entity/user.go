package user_entity

import (
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
)

type User struct {
	ID               uint
	ArticlesUploaded []*article_entity.Article `gorm:"foreignKey:AuthorID,constraint:OnDelete:CASCADE;"`
	ArticlesLiked    []*article_entity.Article `gorm:"many2many:user_likes;"`
	CommentsCreated  []*comment_entity.Comment `gorm:"foreignKey:AuthorID,constraint:OnDelete:SET NULL;"`
}
