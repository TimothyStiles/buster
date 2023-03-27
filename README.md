# buster

## dagger workflows for Go CI/CD

[dagger](https://github.com/dagger/dagger) is a new-ish framework for developing CI/CD workflows in Go, Python, or whatever else they support.

Long story short. I don't like Yaml and I want to be able to define my most common CI/CD processes as a Go package and import them to be run somewhere else. Ultimately most CI/CD platforms will require some YAML but the goal is to cut down on it as much as possible to make CI/CD workflows easy to develop, maintain, and ship.

Here's what I'm thinking

Simple:

1. Run tests for various OSs
2. Report test coverage in PR thread. Push to coverage badge like current workflow.
3. Auto-lint and make changes to be autosubmitted PR to current PR branch
4. New goreleaser like workflow?
5. Using new github blocks show test coverage in PR overview.
6. Unneccesary conversion check