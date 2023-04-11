package github_client

import (
	"context"
	"sync"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

const (
	API_URL = "https://api.github.com/graphql"
)

var clientOnce = sync.Once{}

var client *graphql.Client

func getClient() *graphql.Client {
	if client == nil {
		clientOnce.Do(func() {
			src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.GH_TOKEN})
			httpClient := oauth2.NewClient(context.Background(), src)
			client = graphql.NewClient(API_URL, httpClient)
		})
	}
	return client
}
