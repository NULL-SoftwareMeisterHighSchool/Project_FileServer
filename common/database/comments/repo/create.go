package comment_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
)

func CreateComment(articleID, userID uint, replyTo uint, body string) error {

	comment := comment_entity.Comment{
		AuthorID:  userID,
		ArticleID: articleID,
		Body:      body,
	}

	if replyTo != 0 {
		comment.ReplyCommentID = &replyTo
	}

	tx := database.Comments().Create(&comment)
	return tx.Error
}
