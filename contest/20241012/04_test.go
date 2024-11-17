// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/10/12 23:30
package _0241012

import "testing"

func Test_power(t *testing.T) {
	type args struct {
		base     int
		exponent int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 3}, 8},
		{"2", args{3, 3}, 27},
		{"3", args{5, 3}, 125},
		{"4", args{7, 3}, 343},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := power(tt.args.base, tt.args.exponent); got != tt.want {
				t.Errorf("power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_comb(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 3}, 60},
		{"2", args{5, 2}, 20},
		{"3", args{10, 2}, 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comb(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("comb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfWays(t *testing.T) {
	type args struct {
		n int
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	"001",
		//	args{1, 2, 3},
		//	6,
		//},
		//{
		//	"002",
		//	args{5, 2, 1},
		//	32,
		//},
		{
			"003",
			args{3, 3, 4},
			684,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.n, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
