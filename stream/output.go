package stream

import (
	"fmt"
	"io"
	"net/http"
)

// OutputFromStream reads data from a passed stream and writes it to another stream
func OutputFromStream(r io.Reader, w io.Writer) {
	content, err := io.ReadAll(r)
	if err != nil {
		return
	}
	_, err = w.Write(content)
	if err != nil {
		return
	}
}

// PrintRobotsExist outputs the status of exist file
func PrintRobotsExist(r *http.Response) {
	fmt.Printf("URL : %s\nExist : %t\n", r.Request.URL, ExistRobots(r))
}
