package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	image_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/images/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
)

func DeleteByID(userID, articleID uint) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	storages.Get().
		DeleteImages(
			image_repo.GetURLsByArticleID(articleID),
		)

	return article_repo.DeleteByID(articleID)
}
