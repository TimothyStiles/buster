package tester_test

import (
	"fmt"

	"github.com/TimothyStiles/buster/config"
	"github.com/TimothyStiles/buster/tester"
)

func ExampleRunTests() {
	cfg := config.NewConfig()
	cfg.RootPath = "../"

	_, err := tester.RunTests(cfg, "./tester/mock")

	fmt.Println(err)
}
