package main

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// Set the environment variables
	os.Setenv("GITHUB_GIST_TOKEN", "your_token")
	os.Setenv("GITHUB_GIST_FILENAME", "test.txt")
	os.Setenv("GITHUB_GIST_ID", "")
	os.Setenv("GITHUB_GIST_MESSAGE", "test_message")
	os.Setenv("GITHUB_GIST_COLOR", "blue")
	os.Setenv("GITHUB_GIST_LABEL", "test_label")
	os.Setenv("SHIELD_IO_SCHEMA_VERSION", "1")

	// Call the getEnv function
	token, badge := getEnv()

	// Assert the values
	expectedToken := "your_token"
	if token != expectedToken {
		t.Errorf("Expected token to be %s, but got %s", expectedToken, token)
	}

	expectedFilename := "test.txt"
	if badge.Filename != expectedFilename {
		t.Errorf("Expected filename to be %s, but got %s", expectedFilename, badge.Filename)
	}

	expectedGistID := ""
	if *badge.Gist.ID != expectedGistID {
		t.Errorf("Expected Gist ID to be %s, but got %s", expectedGistID, *badge.Gist.ID)
	}

	expectedMessage := "test_message"
	if badge.Shield.Message != expectedMessage {
		t.Errorf("Expected shield message to be %s, but got %s", expectedMessage, badge.Shield.Message)
	}

	expectedColor := "blue"
	if badge.Shield.Color != expectedColor {
		t.Errorf("Expected shield color to be %s, but got %s", expectedColor, badge.Shield.Color)
	}

	expectedLabel := "test_label"
	if badge.Shield.Label != expectedLabel {
		t.Errorf("Expected shield label to be %s, but got %s", expectedLabel, badge.Shield.Label)
	}

	expectedSchemaVersion := "1"
	if badge.Shield.SchemaVersion != expectedSchemaVersion {
		t.Errorf("Expected shield schema version to be %s, but got %s", expectedSchemaVersion, badge.Shield.SchemaVersion)
	}
}

// func TestCmd(t *testing.T) {
// 	// Set the environment variables
// 	os.Setenv("GITHUB_GIST_TOKEN", "insert_real_token_here")
// 	os.Setenv("GITHUB_GIST_FILENAME", "upsertbadgetest.json")
// 	os.Setenv("GITHUB_GIST_ID", "")
// 	os.Setenv("GITHUB_GIST_MESSAGE", "test_message")
// 	os.Setenv("GITHUB_GIST_COLOR", "blue")
// 	os.Setenv("GITHUB_GIST_LABEL", "test_label")
// 	os.Setenv("SHIELD_IO_SCHEMA_VERSION", "1")

// 	// Call the cmd function
// 	cmd()

// 	// TODO: Add assertions for the expected behavior of the cmd function
// }
