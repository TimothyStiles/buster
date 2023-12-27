package badges

import (
	"encoding/json"
	"errors"
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
	URL      string
	Filename string
	Shield   Shield
	Gist     *github.Gist
}

func (badge *Badge) BuildURL(gist *github.Gist) error {
	urlBase := "https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/"
	owner := gist.Owner.GetLogin()
	if owner == "" {
		return errors.New("gist owner not found")
	}

	id := gist.GetID()
	if id == "" {
		return errors.New("gist id not found")
	}

	filename := badge.Filename
	if filename == "" {
		return errors.New("gist filename not found")
	}

	raw := "raw"
	result, err := url.JoinPath(owner, id, raw, filename)
	if err != nil {
		return err
	}
	badge.URL = urlBase + result

	return nil
}

func (badge *Badge) Upsert(service gists.GistsServiceInterface) error {
	// Marshal the shield into JSON
	shieldJSON, err := json.Marshal(badge.Shield)
	if err != nil {
		return err
	}
	shieldString := string(shieldJSON)

	// check if gist exists
	gistID, err := gists.GetGistID(service, badge.Filename)

	if err != nil {
		return err
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

	err = badge.BuildURL(gist)

	return err
}
