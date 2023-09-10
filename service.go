package githubwrap

import (
	"context"
	"github.com/gofri/go-github-ratelimit/github_ratelimit"
	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
	"net/http"
)

func NewOauth2Client(ctx context.Context, accessToken string) (client *http.Client) {
	if accessToken == "" {
		panic("access token is empty")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	client = oauth2.NewClient(ctx, ts)
	return
}

func NewOauth2ClientLimited(ctx context.Context, accessToken string) (client *http.Client) {
	if accessToken == "" {
		panic("access token is empty")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	client, err := github_ratelimit.NewRateLimitWaiterClient(oauth2.NewClient(ctx, ts).Transport)
	if err != nil {
		panic(err)
	}
	return
}

func NewGithubClient(httpClient *http.Client) (client *github.Client) {
	client = github.NewClient(httpClient)
	return
}

func NewFilesService(client *github.Client, owner, repo string) (wrapper *FilesService) {
	wrapper = &FilesService{
		client: client,
		owner:  owner,
		repo:   repo,
	}
	return
}
