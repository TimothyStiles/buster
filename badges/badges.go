package badges

import (
	"context"
	"encoding/json"
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

func (badge *Badge) Create(client *github.Client) error {
	// create gist
	shieldJSON, err := json.Marshal(badge.Shield)
	if err != nil {
		return err
	}
	shieldString := string(shieldJSON)

	gist := github.Gist{
		Description: github.String(badge.Filename),
		Public:      github.Bool(true),
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(badge.Filename): {
				Content: github.String(shieldString),
			},
		},
	}

	createdGist, _, err := client.Gists.Create(context.Background(), &gist)

	if err != nil {
		return err
	}

	badge.Gist = createdGist
	badge.BuildURL(badge.Gist)
	return nil
}

func (badge *Badge) Update(client *github.Client) error {
	// update gist
	shieldJSON, err := json.Marshal(badge.Shield)
	if err != nil {
		return err
	}
	shieldString := string(shieldJSON)

	gist := github.Gist{
		Description: github.String(badge.Filename),
		Public:      github.Bool(true),
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(badge.Filename): {
				Content: github.String(shieldString),
			},
		},
	}

	updatedGist, _, err := client.Gists.Edit(context.Background(), badge.Gist.GetID(), &gist)

	if err != nil {
		return err
	}

	badge.Gist = updatedGist
	return nil
}

func (badge *Badge) Get(client *github.Client) error {

	var err error
	if badge.Gist.GetID() == "" {
		err = badge.getIDFromDescription(client)
	}

	if err != nil {
		return err
	}
	// read gist
	gist, _, err := client.Gists.Get(context.Background(), badge.Gist.GetID())

	if err != nil {
		return err
	}

	badge.Gist = gist
	return nil
}

func (badge *Badge) getIDFromDescription(client *github.Client) error {

	user, _, err := client.Users.Get(context.Background(), "")

	gists, _, err := client.Gists.List(context.Background(), user.GetLogin(), nil)
	if err != nil {
		return err
	}

	for _, gist := range gists {
		if gist.GetDescription() == badge.Filename {
			badge.Gist = gist
		}
	}
	return nil
}

func (badge *Badge) Upsert(client *github.Client) error {

	err := badge.Get(client)
	if err == nil {
		// update gist
		err = badge.Update(client)
		if err != nil {
			return err
		}

	} else {
		// create gist
		err = badge.Create(client)
		if err != nil {
			return err
		}
	}

	err = badge.BuildURL(badge.Gist)

	if err != nil {
		return err
	}

	return nil
}
