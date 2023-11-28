package address

import (
	"errors"
	"net/url"
)

const (
	protocolSecure  = "https"
	protocolDefault = "http"
	robotsText      = "robots.txt"
)

var (
	ErrInvalidURLArguments error = errors.New("invalid url given")
)

// CreateRobotsAddress creates a robots.txt address from a given url with https.
func CreateHTTPSRobotsAddress(addr string) (string, error) {
	if addr == "" {
		return "", ErrInvalidURLArguments
	}

	u, err := url.Parse(addr)
	if err != nil {
		// TODO: should return a error with a wrapped error
		return "", err
	}

	host := u.Host

	robots := &url.URL{}
	robots.Scheme = protocolSecure
	robots.Host = host
	robots.Path = robotsText

	return robots.String(), nil
}

// CreateHTTPRobotsAddress creates a robots.txt address from a given url with http.
func CreateHTTPRobotsAddress(addr string) (string, error) {
	if addr == "" {
		return "", ErrInvalidURLArguments
	}

	u, err := url.Parse(addr)
	if err != nil {
		return "", err
	}
	host := u.Host

	robots := &url.URL{}
	robots.Scheme = protocolDefault
	robots.Host = host
	robots.Path = robotsText

	return robots.String(), nil
}
