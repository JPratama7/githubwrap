package githubwrap

import "github.com/google/go-github/v56/github"

type FilesService struct {
	client *github.Client
	owner  string
	repo   string
}
