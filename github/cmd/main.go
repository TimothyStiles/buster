package main

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	// Get GitHub Gist token from environment variable
	token := os.Getenv("GITHUB_GIST_TOKEN")

	// Create a real GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Use the client...

	// Create a gist
	filename := "test.txt"
	content := "Hello, World!"

	gist := github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): {
				Filename: &filename,
				Content:  &content,
			},
		}}

	createdGist, _, err := client.Gists.Create(ctx, &gist)
	if err != nil {
		panic(err)
	}

	// Print the URL to the created gist
	println(createdGist.GetHTMLURL())

}
