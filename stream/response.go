package stream

import "net/http"

func ExistRobots(res *http.Response) bool {
	return res.StatusCode == http.StatusOK
}
