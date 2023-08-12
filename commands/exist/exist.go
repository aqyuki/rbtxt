package exist

import (
	"fmt"
	"net/http"

	argu "github.com/aqyuki/rbtxt/args"
	"github.com/aqyuki/rbtxt/stream"
	"github.com/aqyuki/rbtxt/stream/url"
	"github.com/spf13/cobra"
)

func ExistCommandHandler(cmd *cobra.Command, args []string) {
	// Check args counts
	if !argu.CheckLength(args, 1) {
		return
	}

	// Create stream
	var res *http.Response
	targetURL := url.CreateRobotsURLSecure(args[0])
	res, err := stream.CreateStream(targetURL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// Check exist
	exist := stream.ExistRobots(res)
	fmt.Printf("Exist %t\n", exist)
}
