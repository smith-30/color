package color

import (
	"testing"
)

func TestRed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "s",
			},
			want: RED + "s" + RESET,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Red(tt.args.s); got != tt.want {
				t.Errorf("Red() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreen(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "s",
			},
			want: GREEN + "s" + RESET,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Green(tt.args.s); got != tt.want {
				t.Errorf("Green() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedf(t *testing.T) {
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:    "s %v",
				args: []interface{}{1},
			},
			want: RED + "s 1" + RESET,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Redf(tt.args.s, tt.args.args...); got != tt.want {
				t.Errorf("Redf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreenf(t *testing.T) {
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:    "s %v",
				args: []interface{}{1},
			},
			want: GREEN + "s 1" + RESET,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greenf(tt.args.s, tt.args.args...); got != tt.want {
				t.Errorf("Greenf() = %v, want %v", got, tt.want)
			}
		})
	}
}
