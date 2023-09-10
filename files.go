package githubwrap

import (
	"context"
	"errors"
	"github.com/google/go-github/v55/github"
	"io"
)

func (wrp *FilesService) Create(ctx context.Context, path, message string, content []byte) (err error) {
	_, _, err = wrp.client.Repositories.CreateFile(ctx, wrp.owner, wrp.repo, path, &github.RepositoryContentFileOptions{
		Content: content,
		Message: github.String(message),
	})
	return
}

func (wrp *FilesService) GetContent(ctx context.Context, path string) (file *github.RepositoryContent, err error) {
	file, _, _, err = wrp.client.Repositories.GetContents(ctx, wrp.owner, wrp.repo, path, &github.RepositoryContentGetOptions{})
	return
}

func (wrp *FilesService) Download(ctx context.Context, path string) (file io.ReadCloser, err error) {
	file, res, err := wrp.client.Repositories.DownloadContents(ctx, wrp.owner, wrp.repo, path, &github.RepositoryContentGetOptions{})
	if res.StatusCode != 200 {
		return nil, errors.New("Download file failed :" + res.Status)
	}
	return
}

func (wrp *FilesService) Delete(ctx context.Context, path, sha string) (file *github.RepositoryContentResponse, err error) {
	file, _, err = wrp.client.Repositories.DeleteFile(ctx, wrp.owner, wrp.repo, path, &github.RepositoryContentFileOptions{
		Message: github.String("Delete file " + path),
		SHA:     github.String(sha),
	})
	return
}

func (wrp *FilesService) Update(ctx context.Context, path, sha string, content []byte) (file *github.RepositoryContentResponse, err error) {
	file, _, err = wrp.client.Repositories.UpdateFile(ctx, wrp.owner, wrp.repo, path, &github.RepositoryContentFileOptions{
		Message: github.String("Update file " + path),
		SHA:     github.String(sha),
		Content: content,
	})
	return
}

func (wrp *FilesService) GetLink(ctx context.Context, path string) (link string, err error) {
	file, err := wrp.GetContent(ctx, path)
	if err != nil || file.DownloadURL == nil {
		return
	}

	link = *file.DownloadURL
	return
}

func (wrp *FilesService) ChangeRepository(owner, repo string) {
	wrp.owner = owner
	wrp.repo = repo
	return
}
