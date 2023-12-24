package github

import (
	"bytes"
	"net/http"
)

func CreateGist(url, token string) (*http.Response, error) {
	data := []byte(`{"description":"Example of a gist","public":false,"files":{"README.md":{"content":"Hello World"}}}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
