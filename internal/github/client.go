package github

import (
	"context"
	"fmt"
	"io"

	"github.com/google/go-github/v32/github"
	"github.com/sirupsen/logrus"
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

func WorkFlowDispatch(client *github.Client, owner string, repo string, workflow string) error {
	branch, _ := GetDefaultBranch(client, owner, repo)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/example-wf-dispatch.yml/dispatches", owner, repo)

	payload := map[string]interface{}{
		"ref": branch,
	}

	req, err := client.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	res, err := client.Do(context.Background(), req, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}
	defer res.Body.Close()

	logrus.Info(fmt.Printf("ステータス:%s", res.Status))
	body, _ := io.ReadAll(res.Body)
	logrus.Info(fmt.Printf("ステータス:%s", body))

	return nil
}
