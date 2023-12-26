package gists

import (
	"context"

	"github.com/google/go-github/v57/github"
)

// Define an interface that includes the methods you want to mock
type GistsServiceInterface interface {
	Create(ctx context.Context, gist *github.Gist) (*github.Gist, *github.Response, error)
}

// Modify the CreateGist function to accept a GistsServiceInterface parameter instead of a *github.Client
func CreateGist(service GistsServiceInterface, filename, content string) (*github.Gist, error) {
	file := github.GistFile{
		Filename: &filename,
		Content:  &content,
	}

	gist := github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): file,
		},
		Public: github.Bool(false),
	}

	createdGist, _, err := service.Create(context.Background(), &gist)
	if err != nil {
		return nil, err
	}

	return createdGist, nil
}
