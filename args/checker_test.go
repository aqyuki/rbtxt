package args_test

import (
	"testing"

	arg "github.com/aqyuki/rbtxt/args"
)

func TestCheckLength(t *testing.T) {
	type args struct {
		args []string
		len  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Normal - parameter matching",
			args: args{
				args: []string{
					"sample",
				},
				len: 1,
			},
			want: true,
		},
		{
			name: "Normal - Argument number match (no parameter)",
			args: args{
				args: []string{},
				len:  0,
			},
		},
		{
			name: "Normal - parameter-driven discrepancy (few arguments)",
			args: args{
				args: []string{
					"sample",
				},
				len: 2,
			},
			want: false,
		},
		{
			name: "Normal - parameter-driven discrepancy (more arguments)",
			args: args{
				args: []string{
					"sample",
					"sample2",
					"sample3",
				},
				len: 2,
			},
			want: false,
		},
		{
			name: "Normal - argument number match (parameter nil)",
			args: args{
				args: nil,
				len:  0,
			},
			want: true,
		},
		{
			name: "Normal - argument number un match (parameter nil)",
			args: args{
				args: nil,
				len:  1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arg.CheckLength(tt.args.args, tt.args.len); got != tt.want {
				t.Errorf("Case %s fails. got -> %+v , want -> %+v", tt.name, got, tt.want)
			}
		})
	}
}
