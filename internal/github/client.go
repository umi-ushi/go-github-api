package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func NewGitHubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}

func GetDefaultBranch(client *github.Client, owner string, repo string) (string, error) {
	repository, _, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		return "", err
	}
	if repository.DefaultBranch == nil {
		return "", fmt.Errorf("default branch not found")
	}
	return *repository.DefaultBranch, nil
}
