package show

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	argu "github.com/aqyuki/rbtxt/args"
	"github.com/aqyuki/rbtxt/stream"
	"github.com/aqyuki/rbtxt/stream/url"
	"github.com/spf13/cobra"
)

var (
	SaveFilePath string
)

func ShowCommandHandler(cmd *cobra.Command, args []string) {
	if !argu.CheckLength(args, 1) {
		return
	}

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
}
