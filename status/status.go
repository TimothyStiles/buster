/*
Package status is an example of a Dagger package that checks the status of the dagger runtime to see if it is healthy.
*/

package status

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

// Check is a function that checks the status of the dagger runtime to see if it is healthy. It's mostly used as an example.
func Check() error {
	ctx := context.Background()
	fmt.Println("Building with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	fmt.Println("dagger runtime is healthy")

	return nil
}
