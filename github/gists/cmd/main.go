package main

import (
	"context"
	"os"

	"github.com/TimothyStiles/buster/github/gists"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func main() {

	// Get GitHub Gist token from environment variable
	token := os.Getenv("GITHUB_GIST_TOKEN")
	filename := os.Getenv("GITHUB_GIST_FILENAME")
	content := os.Getenv("GITHUB_GIST_CONTENT")

	// Create a real GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Use the client...

	createdGist, err := gists.CreateGist(client.Gists, filename, content)
	if err != nil {
		panic(err)
	}

	// Print the URL to the created gist
	println(createdGist.GetHTMLURL())

}
