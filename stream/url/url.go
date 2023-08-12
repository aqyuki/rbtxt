package url

import "net/url"

const (
	ProtocolSecure  = "https"
	ProtocolDefault = "http"
	RobotsText      = "robots.txt"
)

func CreateRobotsURLSecure(rawURL string) string {
	if rawURL == "" {
		return ""
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	host := u.Host

	robotsURL := &url.URL{}
	robotsURL.Scheme = ProtocolSecure
	robotsURL.Host = host
	robotsURL.Path = RobotsText
	return robotsURL.String()
}
