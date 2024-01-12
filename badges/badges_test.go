package badges

import (
	"testing"

	"github.com/google/go-github/v57/github"
)

// TestBadge_BuildURL tests the BuildURL method of the Badge struct
func TestBuildURL(t *testing.T) {
	// Create a mock gist
	filename := "coverage.json"
	content := "Hello, World!"
	file := github.GistFile{
		Filename: &filename,
		Content:  &content,
	}
	mockGist := github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): file,
		},
		Description: github.String("coverage.json"),
		ID:          github.String("e58f265655ac0acacdd1a38376ccd32a"),
		Owner: &github.User{
			Login: github.String("TimothyStiles"),
		},
	}

	badge := Badge{
		Filename: "coverage.json",
	}

	err := badge.BuildURL(&mockGist)
	if err != nil {
		t.Errorf("BuildURL returned an error: %v", err)
	}

	// Add your assertions here
}
