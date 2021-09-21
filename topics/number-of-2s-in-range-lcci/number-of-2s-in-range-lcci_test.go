// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/20/21 8:34 PM
package number_of_2s_in_range_lcci

import "testing"

func Test_numberOf2sInRange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	"",
		//	args{4},
		//	1,
		//},
		//{
		//	"",
		//	args{20},
		//	3,
		//},
		//{
		//	"",
		//	args{30},
		//	13,
		//},
		{
			"",
			args{120},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOf2sInRange(tt.args.n); got != tt.want {
				t.Errorf("numberOf2sInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
