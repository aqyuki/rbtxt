package stream

import (
	"errors"
	"net/http"

	urul "github.com/aqyuki/rbtxt/stream/url"
)

// CreateStream create a stream to read robots.txt
func CreateStream(url string) (*http.Response, error) {

	if res, err := http.Get(url); err == nil {
		return res, nil
	} else {
		if errors.Is(err, http.ErrNotSupported) {
			res, err := http.Get(urul.CreateRobotsURLDefault(url))
			if err != nil {
				return nil, err
			}
			return res, err
		}
		return nil, err
	}
}
