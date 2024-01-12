package badges

import (
	"errors"
	"net/url"

	"github.com/google/go-github/v57/github"
)

type Shield struct {
	Label         string `json:"label"`
	Message       string `json:"message"`
	SchemaVersion string `json:"schemaVersion"`
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
