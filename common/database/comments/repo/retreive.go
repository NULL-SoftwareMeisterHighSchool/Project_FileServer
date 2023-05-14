package comment_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
)

func GetOwnerByIDAndArticleID(commentID, articleID uint) (uint, error) {
	var authorID uint
	err := database.Comments.
		Where("id = ? AND article_id = ?", commentID, articleID).
		Select("author_id").
		First(&authorID).
		Error

	return authorID, err
}

func GetCommentsByArticleID(articleID uint) ([]*comment_entity.Comment, error) {
	var comments []*comment_entity.Comment

	tx := database.Comments.
		Where("article_id = ?", articleID).
		Find(&comments)
	return comments, tx.Error
}
