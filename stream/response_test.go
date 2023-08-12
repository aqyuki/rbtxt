package stream_test

import (
	"net/http"
	"testing"

	"github.com/aqyuki/rbtxt/stream"
)

func TestExistRobots(t *testing.T) {
	type args struct {
		res *http.Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Normal - exist file",
			args: args{
				res: &http.Response{
					StatusCode: http.StatusOK,
				},
			},
			want: true,
		},
		{
			name: "Normal - not exist file",
			args: args{
				res: &http.Response{
					StatusCode: http.StatusNotFound,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stream.ExistRobots(tt.args.res); got != tt.want {
				t.Errorf("%s failure. got ->  %v, want -> %v", tt.name, got, tt.want)
			}
		})
	}
}
