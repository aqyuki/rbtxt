package url_test

import (
	"testing"

	"github.com/aqyuki/rbtxt/stream/url"
)

func TestCreateRobotsURLSecure(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal - completed url",
			args: args{
				url: "https://www.hoge.com/robots.txt",
			},
			want: "https://www.hoge.com/robots.txt",
		},
		{
			name: "Normal - url given",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want: "https://www.hoge.com/robots.txt",
		},
		{
			name: "Normal - protocol http",
			args: args{
				url: "http://www.hoge.com/robots.txt",
			},
			want: "https://www.hoge.com/robots.txt",
		},
		{
			name: "Normal - protocol http and url",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want: "https://www.hoge.com/robots.txt",
		},
		{
			name: "Error - empty url",
			args: args{
				url: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := url.CreateRobotsURLSecure(tt.args.url)
			if got != tt.want {
				t.Errorf("%s returned an unexpected return value. want -> %s, got -> %s", tt.name, tt.want, got)
			}
		})
	}
}

func TestCreateRobotsURLDefault(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal - completed url",
			args: args{
				url: "https://www.hoge.com/robots.txt",
			},
			want: "http://www.hoge.com/robots.txt",
		},
		{
			name: "Normal - url given",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want: "http://www.hoge.com/robots.txt",
		},
		{
			name: "Normal - protocol http",
			args: args{
				url: "http://www.hoge.com/robots.txt",
			},
			want: "http://www.hoge.com/robots.txt",
		},
		{
			name: "Normal - protocol http and url",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want: "http://www.hoge.com/robots.txt",
		},
		{
			name: "Error - empty url",
			args: args{
				url: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := url.CreateRobotsURLDefault(tt.args.url); got != tt.want {
				t.Errorf("%s returned an unexpected return value. want -> %s, got -> %s", tt.name, tt.want, got)
			}
		})
	}
}
