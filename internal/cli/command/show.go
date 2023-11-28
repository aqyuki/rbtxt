package command

import (
	"github.com/aqyuki/rbtxt/internal/robots"
	"github.com/spf13/cobra"
)

var (
	SaveFilePath string
)

func CreateShowCommand() *cobra.Command {
	var command *cobra.Command = &cobra.Command{
		Use:   "show",
		Short: "show outputs robot.txt",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			robot, err := robots.NewDefaultDownloader().Download(args[0])
			if err != nil {
				return
			}
			robot.DisplaySimple()
		},
	}
	return command
}
