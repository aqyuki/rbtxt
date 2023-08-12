package main

import (
	"os"

	"github.com/aqyuki/rbtxt/commands/root"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "rbtxt",
		Short: "rbtxt can quickly retrieve and check robot.txt",
		Run:   root.RootCommandHandler,
	}
)

func CommandExecute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(-1)
	}
}
