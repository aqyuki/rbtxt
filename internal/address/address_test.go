package address_test

import (
	"testing"

	"github.com/aqyuki/rbtxt/internal/address"
	"github.com/stretchr/testify/assert"
)

func TestCreateHTTPSRobotsAddress(t *testing.T) {
	t.Parallel()

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Normal - completed url",
			args: args{
				url: "https://www.hoge.com/robots.txt",
			},
			want:    "https://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Normal - url given",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want:    "https://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Normal - protocol http",
			args: args{
				url: "http://www.hoge.com/robots.txt",
			},
			want:    "https://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Normal - protocol http and url",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want:    "https://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Error - empty url",
			args: args{
				url: "",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := address.CreateHTTPSRobotsAddress(tt.args.url)
			if tt.wantErr {
				assert.Error(t, err, "error should be returned")
			} else {
				assert.NoError(t, err, "error should not be returned")
			}
			assert.EqualValues(t, tt.want, got, "want and got should be equal")
		})
	}
}

func TestCreateHTTPRobotsAddress(t *testing.T) {
	t.Parallel()

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Normal - completed url",
			args: args{
				url: "https://www.hoge.com/robots.txt",
			},
			want:    "http://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Normal - url given",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want:    "http://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Normal - protocol http",
			args: args{
				url: "http://www.hoge.com/robots.txt",
			},
			want:    "http://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Normal - protocol http and url",
			args: args{
				url: "https://www.hoge.com/top",
			},
			want:    "http://www.hoge.com/robots.txt",
			wantErr: false,
		},
		{
			name: "Error - empty url",
			args: args{
				url: "",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := address.CreateHTTPRobotsAddress(tt.args.url)
			if tt.wantErr {
				assert.Error(t, err, "error should be returned")
			} else {
				assert.NoError(t, err, "error should not be returned")
			}
			assert.EqualValues(t, tt.want, got, "want and got should be equal")
		})
	}
}
