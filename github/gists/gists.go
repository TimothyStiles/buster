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

// CreateGist creates a new gist with the given filename and content to the service token's user's account
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

// ChecksGistsExists gets the names of all gists of the token's user and checks if a gist of that name already exists under the user's account
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
