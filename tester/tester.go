package tester

import (
	"github.com/TimothyStiles/buster/config"
	"github.com/TimothyStiles/buster/runner"
)

// RunTestContainer runs a test container with the given config and path to tests.
func RunTestContainer(cfg *config.Config, path string) (string, error) {

	cfg, client, ctx := runner.DefaultSetup(cfg)

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
