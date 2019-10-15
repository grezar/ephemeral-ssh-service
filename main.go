package main

import (
	"fmt"
	"os"

	"github.com/grezar/ephemeral-ssh-service/cmd"
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
