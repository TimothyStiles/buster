package status_test

import "github.com/TimothyStiles/buster/status"

func ExampleCheck() {
	err := status.Check()

	if err != nil {
		panic(err)
	}
	// Output:
	// Building with Dagger
	// dagger runtime is healthy
}
