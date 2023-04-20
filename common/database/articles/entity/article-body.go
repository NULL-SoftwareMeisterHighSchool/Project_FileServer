package article_entity

type ArticleBody struct {
	ArticleID uint
	Body      []byte `gorm:"type:longtext"`
}
