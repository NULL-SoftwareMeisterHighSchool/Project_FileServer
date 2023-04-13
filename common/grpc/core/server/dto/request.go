package core_server_dto

type GithubUserInfo struct {
	ID          uint
	GithubLogin string
}

// user request from core
type RetrieveContributionReq []*GithubUserInfo
