package robots

import (
	"errors"
	"io"
	"net/http"

	"github.com/aqyuki/rbtxt/internal/address"
)

type Downloader struct {
	useHTTP bool
}

var (
	ErrFailedToDownload = errors.New("failed to download robots.txt")
)

// Download try to download robots.txt from url
func (d *Downloader) Download(url string) (*Robots, error) {
	if res, err := d.DownloadWithHTTPS(url); err == nil {
		return res, nil
	} else if errors.Is(err, http.ErrNotSupported) && d.useHTTP {
		res, err := d.DownloadWithHTTP(url)
		if err != nil {
			return nil, err
		}
		return res, err
	}
	return nil, ErrFailedToDownload
}

// DownloadWithHTTP try to download robots.txt from url with https
func (d *Downloader) DownloadWithHTTPS(url string) (*Robots, error) {

	target, err := address.CreateHTTPSRobotsAddress(url)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := collectBytesFromResponse(res)
	if err != nil {
		return nil, err
	}
	return NewRobots(true, target, body), nil
}

// Download try to download robots.txt from url with http
func (d *Downloader) DownloadWithHTTP(url string) (*Robots, error) {

	target, err := address.CreateHTTPRobotsAddress(url)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := collectBytesFromResponse(res)
	if err != nil {
		return nil, err
	}
	return NewRobots(true, target, body), nil
}

func collectBytesFromResponse(res *http.Response) (string, error) {
	b, err := io.ReadAll(res.Body)
	body := string(b)
	return body, err
}

// NewDefaultDownloader create a downloader with default options
func NewDefaultDownloader() *Downloader {
	return &Downloader{
		useHTTP: true,
	}
}

type Option func(d *Downloader)

// UseHTTP use http to download robots.txt
func UseHTTP() Option {
	return func(d *Downloader) {
		d.useHTTP = true
	}
}

// NewDownloaderWithOption create a downloader with options
func NewDownloaderWithOption(opts ...Option) *Downloader {
	d := &Downloader{}
	for i := range opts {
		opts[i](d)
	}
	return d
}
