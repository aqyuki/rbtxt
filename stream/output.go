package stream

import (
	"fmt"
	"io"
	"net/http"
)

// PrintFromStream outputs the contents written to the stream to standard output
func PrintFromStream(r io.Reader) {
	content, err := io.ReadAll(r)
	if err != nil {
		return
	}
	fmt.Printf("%s", string(content))
}

// PrintRobotsExist outputs the status of exist file
func PrintRobotsExist(r *http.Response) {
	fmt.Printf("URL : %s\nExist : %t\n", r.Request.URL, ExistRobots(r))
}
