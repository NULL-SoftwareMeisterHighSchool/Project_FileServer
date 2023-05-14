package comment_entity

import (
	"time"

	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
)

type Comment struct {
	ID             uint              `gorm:"primary key"`
	AuthorID       uint              `gorm:"primary key;autoIncrement:false"`
	Author         *user_entity.User `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL;"`
	ArticleID      uint
	ReplyCommentID *uint
	Body           string
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	Replies        []*Comment `gorm:"foreignkey:ReplyCommentID"`
}
