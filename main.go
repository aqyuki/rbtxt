package main

import (
	"os"

	"github.com/aqyuki/rbtxt/internal/cli"
)

func main() {
	cli := cli.CreateDefaultCLI()
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
