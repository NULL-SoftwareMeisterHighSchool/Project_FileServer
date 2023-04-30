package article_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	"gorm.io/gorm"
)

func IncreaseViewCount(articleID uint) {
	// might need to seperate the views into table. someday
	database.Articles.
		Where("id = ?", articleID).
		UpdateColumn("views", gorm.Expr("views  + ?", 1))
}
