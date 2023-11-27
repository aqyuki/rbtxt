package command

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aqyuki/rbtxt/stream"
	"github.com/aqyuki/rbtxt/stream/url"
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
			var res *http.Response
			targetURL := url.CreateRobotsURLSecure(args[0])
			res, err := stream.CreateStream(targetURL)
			if err != nil {
				return
			}
			defer res.Body.Close()

			if SaveFilePath != "" {
				absPath, err := filepath.Abs(SaveFilePath)
				if err != nil {
					fmt.Println("Failed to output to file")
					return
				}
				file, err := os.Create(absPath)
				if err != nil {
					fmt.Println("Failed to output to file")
					return
				}
				defer file.Close()
				stream.OutputFromStream(res.Body, file)
				return
			}

			if stream.ExistRobots(res) {
				stream.OutputFromStream(res.Body, os.Stdout)
			} else {
				fmt.Println("Robots.txt is not installed")
			}
		},
	}
	command.Flags().StringVarP(&SaveFilePath, "save", "s", "", "Save to file")
	return command
}
