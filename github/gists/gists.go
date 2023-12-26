package gists

import (
	"context"

	"github.com/google/go-github/v57/github"
)

// Define an interface that includes the methods you want to mock
type GistsServiceInterface interface {
	Create(ctx context.Context, gist *github.Gist) (*github.Gist, *github.Response, error)
	ListAll(ctx context.Context, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
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

// Function to get the names of all gists and check if a gist of that name already exists
func CheckGistExists(service GistsServiceInterface, gistName string) (bool, error) {
	// List all gists
	gists, _, err := service.ListAll(context.Background(), nil)
	if err != nil {
		return false, err
	}

	// Check if a gist with the given name already exists
	for _, gist := range gists {
		if *gist.Description == gistName {
			return true, nil
		}
	}

	return false, nil
}
