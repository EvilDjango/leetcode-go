// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/19/21 3:28 PM
package get_kth_magic_number_lcci

import (
	"fmt"
	"testing"
)

func Test_getKthMagicNumber(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Println(getKthMagicNumber(i))
	}
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{1},
			1,
		},
		{
			"",
			args{2},
			3,
		},
		{
			"",
			args{3},
			5,
		},
		{
			"",
			args{4},
			7,
		},
		{
			"",
			args{5},
			9,
		},
		{
			"",
			args{6},
			15,
		},
		{
			"",
			args{7},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthMagicNumber(tt.args.k); got != tt.want {
				t.Errorf("getKthMagicNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
