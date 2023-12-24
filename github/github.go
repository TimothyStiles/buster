package github

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/go-github/v57/github"
)

type GitHubService struct {
	URL   string
	Token string
}

func (s *GitHubService) CreateGist(gist *github.Gist) (*github.Gist, error) {
	data, err := json.Marshal(gist)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", s.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+s.Token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createdGist github.Gist
	if err := json.NewDecoder(resp.Body).Decode(&createdGist); err != nil {
		return nil, err
	}

	return &createdGist, nil
}
