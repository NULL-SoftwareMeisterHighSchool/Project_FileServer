package comment_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
)

func CreateComment(articleID, userID, mentionUserID, replyCommentID uint, body string) error {

	comment := comment_entity.Comment{
		AuthorID:  userID,
		ArticleID: articleID,
		Body:      body,
	}

	if mentionUserID != 0 {
		comment.MentionUserID = &mentionUserID
	}

	if replyCommentID != 0 {
		comment.ReplyCommentID = &replyCommentID
	}

	tx := database.Comments().Create(&comment)
	return tx.Error
}
