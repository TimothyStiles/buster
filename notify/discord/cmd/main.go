package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TimothyStiles/buster/notify/discord"
)

func main() {

	// get the webhook from the environment variable
	webhook := os.Getenv("DISCORD_WEBHOOK")
	if webhook == "" {
		log.Fatal("no webhook provided")
	}

	// get the message from the environment variable
	message := os.Getenv("DISCORD_MESSAGE")
	if message == "" {
		log.Fatal("no message provided")
	}

	// send the message to discord
	resp, err := discord.Notify(webhook, message)
	if err != nil {
		log.Fatal(err)
	}

	// handle response
	if resp.StatusCode == http.StatusOK {
		// Successful response
		// You can perform further actions here
		fmt.Println("Request successful")
	} else {
		// Unsuccessful response
		// You can handle different status codes here
		fmt.Println("Request failed with status code:", resp.StatusCode)
	}
}
