package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"dagger.io/dagger"
)

// Test is a function that runs a test command in a container
func Test(exec []string) (string, error) {
	ctx := context.Background()

	if len(exec) == 0 {
		return "", fmt.Errorf("no command provided")
	} else if len(exec) == 1 {
		exec = append([]string{"go", "test"}, exec...)
	} else if exec[0] != "go" && exec[1] != "test" {
		exec = append([]string{"go", "test"}, exec...)
	}

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return "", err
	}
	defer client.Close()

	absRepoPath, err := findGoModOrGit()
	if err != nil {
		log.Println(err)
	}
	// get the source code to test
	src := client.Host().Directory(absRepoPath)

	// get the default container image for golang
	golang := client.Container().From("golang:1.21")

	// mount cloned repository into `golang` image
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	// Print the contents of the mounted directory
	golang = golang.WithExec(exec)

	// run the container
	results, err := golang.Stdout(ctx)

	if err != nil {
		return results, err
	}

	return results, nil
}

// this is a helper function to help find the absolute path across different operating systems and environments
func findGoModOrGit() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		goModPath := filepath.Join(dir, "go.mod")
		gitPath := filepath.Join(dir, ".git")
		if _, err := os.Stat(goModPath); err == nil {
			return dir, nil
		} else if _, err := os.Stat(gitPath); err == nil {
			return dir, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("neither go.mod nor .git found")
}
