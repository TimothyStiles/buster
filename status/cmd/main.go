package main

import (
	"fmt"
	"os"

	"github.com/TimothyStiles/buster/status"
)

func main() {
	if err := status.Check(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
