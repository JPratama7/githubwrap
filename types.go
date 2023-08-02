package githubwrap

import "github.com/google/go-github/v53/github"

type Wrapper struct {
	client *github.Client
	owner  string
	repo   string
}
