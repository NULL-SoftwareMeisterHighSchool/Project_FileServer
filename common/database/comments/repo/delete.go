package comment_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
)

func DeleteComment(commentID, articleID uint) error {
	tx := database.Comments().
		Where("id = ? AND article_id = ?", commentID, articleID).
		Delete(&comment_entity.Comment{})
	return tx.Error
}
