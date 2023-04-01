package tester

import (
	"github.com/TimothyStiles/buster/config"
	"github.com/TimothyStiles/buster/runner"
)

// RunTests runs a test container with the given config and path to tests.
func RunTests(cfg *config.Config, path string) (string, error) {

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

	golang = golang.WithExec([]string{"go", "test", "-v", path})

	// run the container
	results, err := golang.Stdout(*ctx)

	return results, err
}
