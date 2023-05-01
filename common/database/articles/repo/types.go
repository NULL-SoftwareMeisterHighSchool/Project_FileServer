package article_repo

type ArticleOrder uint8

const (
	TIME ArticleOrder = iota
	POPULARITY
	RELEVANCE
)

type ArticleTypeOption uint8

const (
	ALL ArticleTypeOption = iota
	GENERAL
	TECH
)
