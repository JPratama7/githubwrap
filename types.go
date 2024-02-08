package githubwrap

import "github.com/google/go-github/v58/github"

type FilesService struct {
	client *github.Client
	owner  string
	repo   string
}
