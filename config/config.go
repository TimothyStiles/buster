package config

import "os"

type Config struct {
	// ConfigFile is the path to the config file.
	ConfigFile string `yaml:"-"`

	// ProjectRoot is the root of the project.
	RootPath string `yaml:"-"`

	// ContainerImage is the image to use for the container.
	ContainerImage string `yaml:"containerImage"`

	// Commands is a list of commands to run.
	Commands []string `yaml:"commands"`

	// LogOutput is the output for the logs.
	LogOutput *os.File `yaml:"-"`
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{}
}

func DefaultConfig() *Config {
	return &Config{
		ContainerImage: "golang:1.20",
		Commands:       []string{"go", "test", "-v", "."},
		RootPath:       ".",
		ConfigFile:     "buster.JSON",
		LogOutput:      nil, // set to os.Stdout for logging
	}
}
