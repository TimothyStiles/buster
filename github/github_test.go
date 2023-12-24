package github

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGist(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	// Call CreateGist with the mock server's URL
	resp, err := CreateGist(server.URL, "dummy-token")
	if err != nil {
		t.Fatalf("CreateGist returned error: %v", err)
	}

	// Check that CreateGist made a successful request
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("CreateGist returned status %d, want %d", resp.StatusCode, http.StatusCreated)
	}
}
