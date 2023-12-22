package discord

import (
	"bytes"
	"net/http"
)

func Notify(webhook string, message string) (*http.Response, error) {
	jsonStr := []byte(`{"content":"` + message + `"}`)
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
	if err != nil {
		// handle error
		return &http.Response{StatusCode: 500}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return resp, err
	}
	defer resp.Body.Close()

	return resp, nil
}
