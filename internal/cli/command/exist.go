package command

import (
	"fmt"

	"github.com/aqyuki/rbtxt/internal/robots"
	"github.com/spf13/cobra"
)

func CreateExistCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "exist",
		Short: "exist checks whether the target has robot.txt",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			robot, err := robots.NewDefaultDownloader().Download(args[0])
			if err != nil {
				fmt.Println("failed to check robots.txt")
			}

			if robot.Exist {
				fmt.Println("robots.txt is exist.")
			} else {
				fmt.Println("robots.txt is not exist.")
			}
		},
	}
}
