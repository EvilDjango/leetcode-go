// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/12/8 11:22
package weekly_contest_427

import "testing"

func Test_maxSubarraySum(t *testing.T) {
	test(t, maxSubarraySum)
}

func Test_maxSubarraySum2(t *testing.T) {
	test(t, maxSubarraySum2)
}

func Test_maxSubarraySum3(t *testing.T) {
	test(t, maxSubarraySum3)
}

func test(t *testing.T, function func(nums []int, k int) int64) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"case001",
			args{
				[]int{1, 2},
				1,
			},
			3,
		},
		{
			"case002",
			args{
				[]int{-1, -2, -3, -4, -5},
				4,
			},
			-10,
		},
		{
			"case003",
			args{
				[]int{-5, 1, 2, -3, 4},
				2,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
