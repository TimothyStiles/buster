package main

import (
	"fmt"
	"os"

	"github.com/TimothyStiles/buster/test"
)

func main() {
	args := os.Args[1:]
	// if len(args) == 0 {
	// 	fmt.Println("Please provide input from the command line.")
	// 	return
	// }

	// testCommand := []string{"./test/mock/mock_test.go"}
	// fmt.Print(args)
	// if !reflect.DeepEqual(os.Args, testCommand) {
	// 	fmt.Println("Please provide input from the command line.")
	// 	return
	// }
	// args = testCommand
	results, _ := test.Test(args)
	fmt.Println(results)
}
