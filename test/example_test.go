package test_test

import (
	"fmt"

	"github.com/TimothyStiles/buster/test"
)

func ExampleTest() {
	// Write your code here
	testCommand := []string{"go", "test", "./test/mock/mock_test.go"}
	results, _ := test.Test(testCommand)

	fmt.Println(results)
	// Output:
	// "ok  \tcommand-line-arguments\t0.001s\n"
}
