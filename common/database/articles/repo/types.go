package article_repo

type ArticleOrder uint8

const (
	TIME ArticleOrder = iota
	VIEWS
	LIKES
)

type ArticleTypeOption uint8

const (
	GENERAL ArticleTypeOption = iota
	TECH
	INTRODUCE
)
