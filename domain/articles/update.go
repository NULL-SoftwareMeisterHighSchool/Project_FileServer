package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
)

func UpdateArticleBody(articleID, userID uint, body []byte) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	return article_repo.UpdateArticleBody(articleID, body)
}

func UpdateArticleTitle(articleID, userID uint, title string) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	return article_repo.UpdateTitleByID(articleID, title)
}

func UpdateArticleVisibility(userID, articleID uint, isPrivate bool) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	return article_repo.UpdateIsPrivateByID(articleID, isPrivate)
}

func ToggleArticleLike(articleID, userID uint) error {
	if err := article_utils.CheckExists(articleID); err != nil {
		return err
	}

	return article_repo.ToggleLike(articleID, userID)
}
