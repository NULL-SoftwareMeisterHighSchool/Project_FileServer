package comment_entity

import (
	"time"

	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
)

type Comment struct {
	ID             uint `gorm:"primary key"`
	ArticleID      uint
	AuthorID       uint              `gorm:"primary key;autoIncrement:false"`
	Author         *user_entity.User `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;"`
	MentionUserID  *uint
	MentionUser    *user_entity.User `gorm:"foreignKey:MentionUserID;constraint:OnDelete:SET NULL;"`
	Body           string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	ReplyCommentID *uint
	Replies        []*Comment `gorm:"foreignkey:ReplyCommentID;constraint:OnDelete:CASCADE"`
}
