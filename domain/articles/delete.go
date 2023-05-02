package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
)

func DeleteByID(userID, articleID uint) error {
	if err := checkSudoAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.DeleteByID(articleID)

	return nil
}
