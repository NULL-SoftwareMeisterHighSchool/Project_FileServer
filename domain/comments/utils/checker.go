package comment_utils

import (
	"errors"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	comment_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/comments/errors"
	"gorm.io/gorm"
)

func CheckPermissionAndExists(commentID, articleID, userID uint) error {
	authorID, err := comment_repo.GetOwnerByIDAndArticleID(commentID, articleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return comment_errors.ErrCommentNotFound
		}
		return err
	}

	if userID != uint(config.ADMIN_ID) && authorID != userID {
		return comment_errors.ErrNoPermission
	}
	return nil
}
