package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
)

func DeleteByID(userID, articleID uint) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	return article_repo.DeleteByID(articleID)
}
