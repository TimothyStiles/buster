package badges

import (
	"context"
	"testing"

	"github.com/TimothyStiles/ditto"
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
func TestUpsert(t *testing.T) {

	shield := Shield{
		Label:         "coverage",
		Message:       "100%",
		SchemaVersion: "1",
		Color:         "green",
	}

	badge := Badge{
		Filename: "buster-coverage-test-call.json",
		Shield:   shield,
	}

	client := github.NewClient(ditto.Client()).WithAuthToken("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")

	err := badge.Create(client)
	if err != nil {
		t.Errorf("Upsert returned an error: %v", err)
	}

	// to check see if gist was created and has an ID
	if badge.Gist == nil {
		t.Errorf("Upsert did not create a gist")
	}

	if badge.Gist.GetID() == "" {
		t.Errorf("Upsert did not create a gist with an ID")
	}

	// remove gist ID from badge so it can be found by filename/description
	badge.Gist.ID = nil
	err = badge.getIDFromDescription(client)
	if err != nil {
		t.Errorf("getIDFromDescription returned an error: %v", err)
	}

	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		t.Errorf("client.Users.Get returned an error: %v", err)
	}

	// to check if gist was created by the authenticated user
	if badge.Gist.Owner.GetLogin() != user.GetLogin() {
		t.Errorf("Upsert did not create a gist owned by the authenticated user")
	}

	// remove gist ID from badge so it can be found by filename/description again
	badge.Gist.ID = nil

	// test get
	err = badge.Get(client)
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}

	badgeDescription := badge.Gist.GetDescription()
	if badgeDescription != "buster-coverage-test-call.json" {
		t.Errorf("Get did not return the correct gist")
	}

	// test upsert
	err = badge.Upsert(client)
	if err != nil {
		t.Errorf("Upsert returned an error: %v", err)
	}
	// Add your assertions here
}
