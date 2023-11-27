package command

import (
	"net/http"

	"github.com/aqyuki/rbtxt/stream"
	"github.com/aqyuki/rbtxt/stream/url"
	"github.com/spf13/cobra"
)

func CreateExistCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "exist",
		Short: "exist checks whether the target has robot.txt",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			// Create stream
			var res *http.Response
			targetURL := url.CreateRobotsURLSecure(args[0])
			res, err := stream.CreateStream(targetURL)
			if err != nil {
				return
			}
			defer res.Body.Close()

			// Check exist
			stream.PrintRobotsExist(res)
		},
	}
}
