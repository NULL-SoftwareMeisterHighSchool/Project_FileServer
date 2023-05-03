package comment_entity

import "time"

type Comment struct {
	ID             uint `gorm:"primary key"`
	AuthorID       uint `gorm:"primary key;autoIncrement:false"`
	ArticleID      uint
	ReplyCommentID *uint
	Body           string
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	Replies        []*Comment `gorm:"foreignkey:ReplyCommentID"`
}
