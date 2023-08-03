package githubwrap

import (
	"context"
	"github.com/google/go-github/v53/github"
)

func (wrp *Wrapper) UploadFile(ctx context.Context, path, message string, content []byte) (err error) {
	_, _, err = wrp.client.Repositories.CreateFile(ctx, wrp.owner, wrp.repo, path, &github.RepositoryContentFileOptions{
		Content: content,
		Message: github.String(message),
	})
	return
}

func (wrp *Wrapper) GetLink(ctx context.Context, path string) (link string, err error) {
	file, _, _, err := wrp.client.Repositories.GetContents(ctx, wrp.owner, wrp.repo, path, &github.RepositoryContentGetOptions{})
	if err != nil || file.DownloadURL == nil {
		return
	}

	link = *file.DownloadURL
	return
}
