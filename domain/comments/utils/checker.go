package comment_utils

import (
	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	comment_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/comments/errors"
)

func CommentExistsOrNotFound(commentID, articleID uint) error {
	if comment_repo.ExistsByIDAndArticleID(commentID, articleID) {
		return comment_errors.ErrCommentNotFound
	}
	return nil
}
