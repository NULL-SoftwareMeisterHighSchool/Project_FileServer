package core_dto

type ArticleInfo struct {
	IsPublic bool
	AuthorID uint
}

type UserInfo struct {
	ID          uint
	GithubLogin string
}
