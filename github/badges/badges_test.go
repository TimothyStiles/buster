package badges

import (
	"net/http"
	"testing"

	"github.com/TimothyStiles/buster/github/gists"
	"github.com/google/go-github/v57/github"
	"github.com/stretchr/testify/mock"
)

func TestUpsert(t *testing.T) {
	// Create a mock GistsServiceInterface
	mockGists := new(gists.MockGistsService)

	// Create a mock response
	mockResponse := &github.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// Create a test badge
	testBadge := Badge{
		Filename: "test.txt",
		Shield: Shield{
			Label:         "test",
			Message:       "test",
			Color:         "blue",
			SchemaVersion: 1,
		},
	}

	// Create a mock gist
	filename := "test.txt"
	content := "Hello, World!"
	file := github.GistFile{
		Filename: &filename,
		Content:  &content,
	}
	mockGist := github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): file,
		},
		Description: github.String("test.txt"),
		ID:          github.String("test"),
	}

	// Expect the ListAll method to be called once with any arguments, and return a slice containing the mock gist, the mock response and no error
	mockGists.On("ListAll", mock.Anything, mock.Anything).Return([]*github.Gist{&mockGist}, mockResponse, nil)
	mockGists.On("Edit", mock.Anything, mock.Anything, mock.Anything).Return(&mockGist, mockResponse, nil)
	mockGists.On("Create", mock.Anything, mock.Anything).Return(&mockGist, mockResponse, nil)

	_, err := Upsert(mockGists, testBadge)
	if err != nil {
		t.Errorf("Upsert returned an error: %v", err)
	}

	// Add your assertions here to verify the behavior of the Upsert function
	// For example, you can check if the mockService's Upsert method was called with the correct arguments
	// You can also check if the returned gist matches your expectations
}
