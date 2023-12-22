package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Test is a function that runs a test command in a container
func Test(cmdString []string) (string, error) {

	if len(cmdString) == 0 {
		return "", fmt.Errorf("no command provided")
	} else if len(cmdString) == 1 {
		cmdString = append([]string{"go", "test"}, cmdString...)
	} else if cmdString[0] != "go" && cmdString[1] != "test" {
		cmdString = append([]string{"go", "test"}, cmdString...)
	}

	absRepoPath, err := findGoModOrGit()

	if err != nil {
		log.Println(err)
	}

	// change the working directory to the absolute path of the repo
	if err := os.Chdir(absRepoPath); err != nil {
		log.Println(err)
	}

	// execute go test
	cmd := exec.Command(cmdString[0], cmdString[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}

	return string(output), nil

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
