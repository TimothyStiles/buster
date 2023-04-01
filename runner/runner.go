package runner

import (
	"context"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/TimothyStiles/buster/config"
)

// DefaultSetup is a constructor for a new Runner with default values.
func DefaultSetup(cfg *config.Config) (*config.Config, *dagger.Client, *context.Context, error) {

	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}

	defaultConfigTemplate := config.DefaultConfig()

	if cfg.RootPath == "" {
		cfg.RootPath = defaultConfigTemplate.RootPath
	}

	// get absolute path to the project root
	absPath, err := filepath.Abs(cfg.RootPath)
	if err != nil {
		return nil, nil, nil, err
	}
	cfg.RootPath = absPath

	if cfg.ContainerImage == "" {
		cfg.ContainerImage = defaultConfigTemplate.ContainerImage
	}

	return cfg, client, &ctx, err
}
