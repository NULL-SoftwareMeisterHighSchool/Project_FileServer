package core_dto

// article response from core
type ArticleInfoRes struct {
	IsPublic bool
	AuthorID uint
}

// user response from core
type UserInfoRes struct {
	ID          uint
	GithubLogin string
}
