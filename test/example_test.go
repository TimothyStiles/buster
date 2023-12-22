package test_test

import (
	"fmt"
	"strings"

	"github.com/TimothyStiles/buster/test"
)

func ExampleTest() {
	// Write your code here
	testCommand := []string{"go", "test", "./test/mock/mock_test.go"}
	results, _ := test.Test(testCommand)

	// split the results on whitespace and remove empty strings
	resultsSlice := strings.Fields(results)

	// print the first two words of the results since the last will be variable based on how fast the tests run
	result := fmt.Sprintf("%s %s", resultsSlice[0], resultsSlice[1])
	fmt.Println(result)
	// Output: ok command-line-arguments
}
