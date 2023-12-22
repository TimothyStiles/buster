package discord

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotify(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request payload
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		expectedBody := `{"content":"test message"}`
		if body != expectedBody {
			t.Errorf("Unexpected request payload. Got: %s, want: %s", body, expectedBody)
		}

		// Send a mock response
		w.WriteHeader(http.StatusOK)
	}))

	// Make sure to close the server at the end of the test
	defer server.Close()

	// Call the Notify function with the mock server URL
	resp, error := Notify(server.URL, "test message")

	// Check if the request returned an error
	if error != nil {
		t.Errorf("Notify() = %v, want nil", error)
	}

	// Add assertions for the expected behavior
	// ...
	// handle response
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
