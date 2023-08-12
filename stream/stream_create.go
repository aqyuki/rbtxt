package stream

import "net/http"

// CreateStream create a stream to read robots.txt
func CreateStream(url string) (*http.Response, error) {
	return http.Get(url)
}
