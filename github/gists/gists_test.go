package gists

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-github/v57/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Create a struct that embeds testify's mock.Mock
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

func TestCreateGist(t *testing.T) {
	// Create a mock gist
	filename := "test.txt"
	content := "Hello, World!"
	file := github.GistFile{
		Filename: &filename,
		Content:  &content,
	}
	gist := github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): file,
		},
	}

	// Create a new MockGistsService
	mockGists := new(MockGistsService)

	// Create a mock response
	mockResponse := &github.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// Expect the Create method to be called once with any arguments, and return the mock gist, mock response and no error
	mockGists.On("Create", mock.Anything, mock.Anything).Return(&gist, mockResponse, nil)

	// Call the CreateGist function with the mock service
	createdGist, err := CreateGist(mockGists, filename, content)

	// Assert that the mock conditions were met
	mockGists.AssertExpectations(t)

	// Assert that the returned gist is as expected and there was no error
	assert.NoError(t, err)
	assert.Equal(t, &gist, createdGist)
}
func TestCheckGistExists(t *testing.T) {
	// Create a mock GistsService
	mockGists := new(MockGistsService)

	// Create a list of mock gists
	mockGistsList := []*github.Gist{
		{
			Description: github.String("gist1"),
		},
		{
			Description: github.String("gist2"),
		},
		{
			Description: github.String("gist3"),
		},
	}

	// Create a mock response
	mockResponse := &github.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// Mock the ListAll method to return the mock gists list
	mockGists.On("ListAll", mock.Anything, mock.Anything).Return(mockGistsList, mockResponse, nil)

	// Test case: Gist exists
	exists, err := CheckGistExists(mockGists, "gist2")
	assert.NoError(t, err)
	assert.True(t, exists)

	// Test case: Gist does not exist
	exists, err = CheckGistExists(mockGists, "gist4")
	assert.NoError(t, err)
	assert.False(t, exists)

	// Assert that the mock conditions were met
	mockGists.AssertExpectations(t)
}
