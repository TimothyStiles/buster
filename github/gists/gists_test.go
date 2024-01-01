package gists

import (
	"net/http"
	"testing"

	"github.com/google/go-github/v57/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Create a struct that embeds testify's mock.Mock

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
func TestGetGistID(t *testing.T) {
	// Create a mock GistsService
	mockGists := new(MockGistsService)

	// Create a list of mock gists
	mockGistsList := []*github.Gist{
		{
			ID:          github.String("gist1"),
			Description: github.String("gist1"),
		},
		{
			ID:          github.String("gist2"),
			Description: github.String("gist2"),
		},
		{
			ID:          github.String("gist3"),
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
	gistID, err := GetGistID(mockGists, "gist2")
	assert.NoError(t, err)
	assert.Equal(t, "gist2", gistID)

	// Test case: Gist does not exist
	gistID, err = GetGistID(mockGists, "gist4")
	assert.NoError(t, err)
	assert.Equal(t, "", gistID)

	// Assert that the mock conditions were met
	mockGists.AssertExpectations(t)
}
func TestEditGistFile(t *testing.T) {
	// Create a mock GistsService
	mockGists := new(MockGistsService)

	// Define the input parameters
	filename := "test.txt"
	content := "Hello, World!"
	gistID := "gist123"

	// Create a mock response
	mockResponse := &github.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// Expect the Edit method to be called once with the correct arguments, and return the mock gist and mock response
	mockGists.On("Edit", mock.Anything, gistID, mock.Anything).Return(&github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): {
				Content: &content,
			},
		},
	}, mockResponse, nil)

	// Call the EditGist function with the mock service
	editedGist, err := EditGistFile(mockGists, filename, content, gistID)

	// Assert that the mock conditions were met
	mockGists.AssertExpectations(t)

	// Assert that the returned gist is as expected and there was no error
	assert.NoError(t, err)
	assert.Equal(t, &github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): {
				Content: &content,
			},
		},
	}, editedGist)
}
