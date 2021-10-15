// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/15/21 7:46 AM
package maximum_lcci

import "testing"

func Test_maximum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{5, 3},
			5,
		},
		{
			"",
			args{6, 0},
			6,
		},
		{
			"",
			args{-5, -3},
			-3,
		},
		{
			"",
			args{-5, -5},
			-5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
