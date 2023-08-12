package stream

import (
	"fmt"
	"io"
)

// PrintFromStream outputs the contents written to the stream to standard output
func PrintFromStream(r io.Reader) {
	content, err := io.ReadAll(r)
	if err != nil {
		return
	}
	fmt.Printf("%s", string(content))
}
