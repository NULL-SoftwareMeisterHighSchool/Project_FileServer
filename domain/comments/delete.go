package comments

import (
	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
	comment_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/comments/utils"
)

func DeleteComment(commentID, userID, articleID uint) error {

	if err := article_utils.CheckPrivateAndExists(userID, articleID); err != nil {
		return err
	}
	if err := comment_utils.CommentExistsOrNotFound(commentID, articleID); err != nil {
		return err
	}

	return comment_repo.DeleteComment(commentID, articleID)
}
