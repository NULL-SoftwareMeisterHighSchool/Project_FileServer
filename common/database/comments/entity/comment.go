package comment_entity

import "time"

type Comment struct {
	ID             uint
	AuthorID       uint
	ArticleID      uint
	ReplyCommentID *uint
	Body           string
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	Replies        []*Comment `gorm:"foreignkey:ReplyCommentID"`
}
