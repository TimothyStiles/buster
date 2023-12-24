package lint

import (
	"fmt"

	"github.com/TimothyStiles/buster/config"
	"github.com/TimothyStiles/buster/runner"
)

// RunLint runs a containerized GolangCI linter on a Go project.
func RunLint(cfg *config.Config, path string) (string, error) {

	// set default container image before DefaultSetup to avoid default Golang image
	if cfg.ContainerImage == "" {
		cfg.ContainerImage = "golangci/golangci-lint:v1.52.2-alpine"
	}

	cfg, client, ctx, err := runner.DefaultSetup(cfg)
	if err != nil {
		return "", err
	}

	// get reference to the local project
	src := client.Host().Directory(cfg.RootPath)

	// get the default container image for testing (golang)
	golang := client.Container().From(cfg.ContainerImage)

	// mount cloned repository into `golang` image
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	fileTree, err := golang.Directory("/src").Entries(*ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(fileTree)

	golang = golang.WithExec([]string{"golangci-lint", "run", path})

	// run the container
	results, err := golang.Stdout(*ctx)

	return results, err
}
