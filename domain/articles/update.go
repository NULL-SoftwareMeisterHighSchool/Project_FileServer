package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

func UpdateArticleBody(articleID, userID uint, body []byte) error {

	if err := checkPermissionAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.UpdateArticleBodyAndImages(articleID, body, util.ExtractImageURL(body))
	return nil
}

func UpdateArticleTitle(articleID, userID uint, title string) error {
	if err := checkPermissionAndExists(userID, articleID); err != nil {
		return err
	}

	return nil
}
