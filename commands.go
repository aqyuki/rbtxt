package main

import (
	"os"

	"github.com/aqyuki/rbtxt/commands/exist"
	"github.com/aqyuki/rbtxt/commands/root"
	"github.com/aqyuki/rbtxt/commands/show"
	"github.com/aqyuki/rbtxt/commands/version"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "rbtxt",
		Short: "rbtxt can quickly retrieve and check robot.txt",
		Run:   root.RootCommandHandler,
	}
	ExistCmd = &cobra.Command{
		Use:   "exist",
		Short: "Check if robots.txt exists",
		Run:   exist.ExistCommandHandler,
	}
	ShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show robots.txt if target host has robots.txt",
		Run:   show.ShowCommandHandler,
	}
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run:   version.VersionCommandHandler,
	}
)

func init() {
	RootCmd.AddCommand(ExistCmd)
	RootCmd.AddCommand(ShowCmd)
	RootCmd.AddCommand(VersionCmd)

	ShowCmd.Flags().StringVarP(&show.SaveFilePath, "save", "s", "", "Output to file instead of to standard output")
}

func CommandExecute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(-1)
	}
}
