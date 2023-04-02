package runner_test

import (
	"fmt"

	"github.com/TimothyStiles/buster/config"
	"github.com/TimothyStiles/buster/runner"
)

func ExampleDefaultSetup() {
	// first let's create a new config using the default setup
	cfg := config.DefaultConfig()
	cfg.ContainerImage = "golang:1.20.2" // overide the default container image to most specific version of golang

	defer cfg.LogOutput.Close()

	// now we can use the default setup to create the default setup to run a cicd pipeline
	cfg, client, ctx, err := runner.DefaultSetup(cfg)
	defer client.Close()

	// get the default container image for testing (golang)
	golang := client.Container().From(cfg.ContainerImage)

	golang = golang.WithExec([]string{"go", "version"})

	// run the container
	results, _ := golang.Stdout(*ctx)

	if err == nil {
		fmt.Println(results)
	}

	// Output: go version go1.20.2 linux/amd64
}
