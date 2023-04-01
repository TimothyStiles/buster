package tester

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"dagger.io/dagger"
)

func RunTests(path string) error {
	// Construct the command to run the tests
	cmd := exec.Command("go", "test", "-v", path)

	// Set the environment variables for the command
	cmd.Env = append(os.Environ(), "GO111MODULE=on")

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running tests: %v\n%s", err, output)
	}

	// Print the output of the command
	fmt.Println(string(output))

	return nil
}

func RunTestContainer(path string) error {

	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// get absolute path to the project root
	absPath, err := filepath.Abs("../")
	/* Assuming that this is the root of all my problems */
	src := client.Host().Directory(absPath) // get reference to the local project

	golang := client.Container().From("golang:1.20")

	// mount cloned repository into `golang` image
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	// cd into the mounted directory and run the tests
	workingDir, err := golang.Workdir(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(workingDir)

	// run ls to see what's in the mounted directory
	golang = golang.WithExec([]string{"ls"})
	results, err := golang.Stdout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	golang = golang.WithExec([]string{"go", "test", "-v", path})

	// run the container

	results, err = golang.Stdout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	return nil
}
