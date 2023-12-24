package github

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/v57/github"
)

func TestCreateGist(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	// Create a new GitHubService using the mock server URL
	service := GitHubService{
		URL:   server.URL,
		Token: "dummy-token",
	}

	// Create a new gist
	gist := &github.Gist{
		Description: github.String("dummy description"),
		Public:      github.Bool(false),
		Files: map[github.GistFilename]github.GistFile{
			"dummy.txt": {
				Content: github.String("dummy content"),
			},
		}}

	// Call the CreateGist function with the mock server URL
	createdGist, err := service.CreateGist(gist)
	if err != nil {
		t.Errorf("CreateGist() = %v, want nil", err)
	}

	// Check if the created gist is not nil
	if createdGist == nil {
		t.Errorf("CreateGist() = nil, want gist")
	}

}
