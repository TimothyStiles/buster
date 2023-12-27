package gists

import (
	"context"

	"github.com/google/go-github/v57/github"
	"github.com/stretchr/testify/mock"
)

type MockGistsService struct {
	mock.Mock
}

// Implement the Create method
func (m *MockGistsService) Create(ctx context.Context, gist *github.Gist) (*github.Gist, *github.Response, error) {
	args := m.Called(ctx, gist)
	return args.Get(0).(*github.Gist), args.Get(1).(*github.Response), args.Error(2)
}

// Implement the ListAll method
func (m *MockGistsService) ListAll(ctx context.Context, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).([]*github.Gist), args.Get(1).(*github.Response), args.Error(2)
}

// Implement the Edit method
func (m *MockGistsService) Edit(ctx context.Context, gistID string, gist *github.Gist) (*github.Gist, *github.Response, error) {
	args := m.Called(ctx, gistID, gist)
	return args.Get(0).(*github.Gist), args.Get(1).(*github.Response), args.Error(2)
}

// Define an interface that includes the methods you want to mock
type GistsServiceInterface interface {
	Create(ctx context.Context, gist *github.Gist) (*github.Gist, *github.Response, error)
	ListAll(ctx context.Context, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
	Edit(ctx context.Context, gistID string, gist *github.Gist) (*github.Gist, *github.Response, error)
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

// GetGistID gets the ID of the gist with the given name under the token's user's account if it exists if not it returns an empty string
func GetGistID(service GistsServiceInterface, gistName string) (string, error) {
	// List all gists
	gists, _, err := service.ListAll(context.Background(), nil)
	if err != nil {
		return "", err
	}

	// Check if a gist with the given name already exists
	for _, gist := range gists {
		if *gist.Description == gistName {
			return *gist.ID, nil
		}
	}

	return "", nil
}

// EditGistfile edits a file of a gist if it exists under the token's user's account
func EditGistFile(service GistsServiceInterface, filename, content, gistID string) (*github.Gist, error) {
	// Get the gist by ID
	gist, _, err := service.Edit(context.Background(), gistID, &github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): {
				Content: &content,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return gist, nil
}
