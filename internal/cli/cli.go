package cli

import (
	"github.com/aqyuki/rbtxt/internal/cli/command"
	"github.com/spf13/cobra"
)

type CLI struct {
	parent *cobra.Command
}

func (c *CLI) Execute() error {
	return c.parent.Execute()
}

// CreateDefaultCLI is a function that creates default CLI.
func CreateDefaultCLI() *CLI {
	cmd := &CLI{
		parent: &cobra.Command{
			Use:   "rbtxt",
			Short: "rbtxt can quickly retrieve and check robot.txt",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Root().Help()
			},
		},
	}

	registerSubCommands(cmd.parent)
	return cmd
}

// registerSubCommands is a function that registers subcommands.
func registerSubCommands(parent *cobra.Command) *cobra.Command {
	// If parent command is nil, it is a most likely bug. So exit by panic and show stack trace.
	if parent == nil {
		panic("failed to register sub-commands because parent command is nil")
	}

	parent.AddCommand(command.CreateExistCommand())
	parent.AddCommand(command.CreateShowCommand())
	parent.AddCommand(command.CreateVersionCommand())
	return parent
}
