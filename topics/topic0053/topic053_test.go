// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/10/21 5:14 PM
package topic0053

import "testing"

func Test_maxInterval(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{-1, 2, -2}},
			2,
		},
		{
			"",
			args{[]int{1, 2, -1, -3, 5, -7, 6, 2, -3}},
			8,
		},
		{
			"",
			args{[]int{1, 2, -1, -3, 0, -7, 1, 2, -3}},
			3,
		},
		{
			"",
			args{[]int{-1, -2, -3}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("topic0053() = %v, want %v", got, tt.want)
			}
		})
	}
}
