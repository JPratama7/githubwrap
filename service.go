package githubwrap

import (
	"context"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
	"net/http"
)

func NewOauth2Client(ctx context.Context, accessToken string) (client *http.Client) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	client = oauth2.NewClient(ctx, ts)
	return
}

func NewGithubClient(httpClient *http.Client) (client *github.Client) {
	client = github.NewClient(httpClient)
	return
}

func NewWrapper(client *github.Client, owner, repo string) (wrapper *Wrapper) {
	wrapper = &Wrapper{
		client: client,
		owner:  owner,
		repo:   repo,
	}
	return
}
