// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/12/22 11:41
package weekly_contest_429

import "testing"

func Test_maxDistinctElements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"001",
			args{[]int{1, 2, 2, 3, 3, 4}, 2},
			6,
		},
		{
			"002",
			args{[]int{4, 4, 4, 4}, 1},
			3,
		},
		{
			"003",
			args{[]int{9, 5, 5}, 0},
			2,
		},
		{
			"004",
			args{[]int{1, 1, 2, 2}, 0},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistinctElements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxDistinctElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
