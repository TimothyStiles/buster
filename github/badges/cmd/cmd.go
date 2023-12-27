package main

import (
	"context"
	"os"

	"github.com/TimothyStiles/buster/github/badges"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func cmd() {

	// Get GitHub Gist token from environment variable
	token, badge := getEnv()

	// Create a real GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Upsert the badge
	err := badge.Upsert(client.Gists)

	if err != nil {
		panic(err)
	}

	// Print the url to the badge
	println(badge.URL)
}

func getEnv() (string, badges.Badge) {
	// Get GitHub Gist token from environment variable
	token := os.Getenv("GITHUB_GIST_TOKEN")
	gistId := os.Getenv("GITHUB_GIST_ID")
	filename := os.Getenv("GITHUB_GIST_FILENAME")
	message := os.Getenv("GITHUB_GIST_MESSAGE")
	color := os.Getenv("GITHUB_GIST_COLOR")
	label := os.Getenv("GITHUB_GIST_LABEL")

	schemaVersion := os.Getenv("SHIELD_IO_SCHEMA_VERSION")

	// create a new badge
	badge := badges.Badge{
		Filename: filename,
		Shield: badges.Shield{
			Label:         label,
			Message:       message,
			Color:         color,
			SchemaVersion: schemaVersion,
		},
	}

	badge.Gist = &github.Gist{
		ID: github.String(gistId),
	}

	return token, badge

}
