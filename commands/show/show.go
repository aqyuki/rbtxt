package show

import (
	"fmt"
	"net/http"

	argu "github.com/aqyuki/rbtxt/args"
	"github.com/aqyuki/rbtxt/stream"
	"github.com/aqyuki/rbtxt/stream/url"
	"github.com/spf13/cobra"
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

	if stream.ExistRobots(res) {
		stream.PrintFromStream(res.Body)
	} else {
		fmt.Println("Robots.txt is not installed")
	}
}
