package robots

import "fmt"

// Robots is a struct that contains the result of robots.txt
type Robots struct {
	Exist bool
	URL   string
	body  string
}

func (r *Robots) Body() string {
	return r.body
}

func (r *Robots) DisplaySimple() {
	if r.Exist {
		fmt.Printf("robots.txt is exist.\nURL : %s\nBody : \n%s\n", r.URL, r.body)
	} else {
		fmt.Println("robots.txt is not exist.")
	}
}

// NewRobots returns a new Robots
func NewRobots(exist bool, url, body string) *Robots {
	return &Robots{
		Exist: exist,
		URL:   url,
		body:  body,
	}
}
