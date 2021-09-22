// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/22/21 12:26 PM
package calculator_lcci

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"1+2 * 3 "},
			7,
		},
		{
			"",
			args{"1+ 2 * 3/7 "},
			1,
		},
		{
			"",
			args{"1* 4-2 * 3 "},
			-2,
		},
		{
			"",
			args{"1+6 * 5/3 "},
			11,
		},
		{
			"",
			args{"3+2*2"},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
