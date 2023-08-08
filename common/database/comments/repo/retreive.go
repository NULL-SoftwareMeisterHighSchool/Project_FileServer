package comment_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
)

type CommentWithReplyCount struct {
	comment_entity.Comment
	ReplyCount uint
}

func GetOwnerByIDAndArticleID(commentID, articleID uint) (uint, error) {
	var authorID uint
	err := database.Comments().
		Where("id = ? AND article_id = ?", commentID, articleID).
		Select("author_id").
		First(&authorID).
		Error

	return authorID, err
}

func GetCommentsByArticleID(articleID uint) ([]*CommentWithReplyCount, error) {
	var comments []*CommentWithReplyCount

	tx := database.Comments().
		Where("article_id = ?", articleID).
		Select("comments.*, (?) AS reply_count",
			database.Comments().
				Where("reply_comment_id = comments.id").
				Select("COUNT(*)"),
		).
		Omit("reply_comment_id", "article_id").
		Find(&comments)
	return comments, tx.Error
}

func GetRepliesByCommentID(commentID uint) ([]*comment_entity.Comment, error) {
	var comments []*comment_entity.Comment

	tx := database.Comments().
		Where("reply_comment_id = ?", commentID).
		Omit("reply_comment_id", "article_id").
		Find(&comments)
	return comments, tx.Error
}

func GetIsReplyByCommentID(commentID uint) bool {
	var isReply bool

	database.Comments().
		Where("id = ?", commentID).
		Select("reply_comment_id != 0").
		Find(&isReply)

	return isReply
}
