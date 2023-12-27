package badges

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/TimothyStiles/buster/github/gists"
	"github.com/google/go-github/v57/github"
)

type Shield struct {
	Label         string `json:"label"`
	Message       string `json:"message"`
	SchemaVersion int    `json:"schemaVersion"`
	Color         string `json:"color"`
}

type Badge struct {
	URL      url.URL
	Filename string
	Shield   Shield
	Gist     *github.Gist
}

func Upsert(service gists.GistsServiceInterface, badge Badge) (*github.Gist, error) {

	// Marshal the shield into JSON
	shieldJSON, err := json.Marshal(badge.Shield)
	if err != nil {
		return nil, err
	}
	shieldString := string(shieldJSON)
	fmt.Println(shieldString)

	// check if gist exists
	gistID, err := gists.GetGistID(service, badge.Filename)
	fmt.Println(gistID)

	if err != nil {
		return nil, err
	}

	var gistExists bool
	if gistID != "" {
		gistExists = true
	}

	var gist *github.Gist
	if gistExists {
		gist, err = gists.EditGistFile(service, badge.Filename, shieldString, gistID)
	} else {
		gist, err = gists.CreateGist(service, badge.Filename, shieldString)
	}
	return gist, err
}
