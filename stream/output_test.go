package stream_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/aqyuki/rbtxt/stream"
)

func PickStdout(t *testing.T, fnc func()) string {
	t.Helper()
	backup := os.Stdout
	defer func() {
		os.Stdout = backup
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w
	fnc()
	w.Close()
	var buffer bytes.Buffer
	if _, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("fail read buf: %v", err)
	}
	s := buffer.String()
	return s
}

func TestPrintFromStream(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal",
			args: args{
				r: bytes.NewBufferString("Sample"),
			},
			want: "Sample",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickStdout(t, func() { stream.PrintFromStream(tt.args.r) }); got != tt.want {
				t.Errorf("Case %s Failure. got -> %s, want -> %s", tt.name, got, tt.want)
			}
		})
	}
}
