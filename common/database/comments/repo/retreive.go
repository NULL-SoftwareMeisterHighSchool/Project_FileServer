package comment_repo

import "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"

func ExistsByIDAndArticleID(commentID, articleID uint) bool {
	var exists bool
	database.Comments.
		Where("id = ? AND article_id = ?", commentID, articleID).
		Select("COUNT(*) > 0").
		Find(&exists)

	return exists
}
