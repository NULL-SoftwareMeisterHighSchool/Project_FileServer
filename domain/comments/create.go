package comments

import (
	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
	comment_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/comments/errors"
)

func CreateComment(articleID, userID uint, replyTo uint, mentionUserID uint, body string) error {
	if _, err := article_utils.CheckPrivateAndExists(userID, articleID); err != nil {
		return err
	}

	if replyTo != 0 {
		isReply := comment_repo.GetIsReplyByCommentID(replyTo)
		if isReply {
			return comment_errors.ErrReplyingToReply
		}
	}

	return comment_repo.CreateComment(articleID, userID, mentionUserID, replyTo, body)
}
