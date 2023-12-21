package status

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

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
