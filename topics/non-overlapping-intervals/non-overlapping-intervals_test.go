// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/16/21 5:36 PM
package non_overlapping_intervals

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	test(t, eraseOverlapIntervals)
}

func Test_eraseOverlapIntervals2(t *testing.T) {
	test(t, eraseOverlapIntervals2)
}

func Test_eraseOverlapIntervals3(t *testing.T) {
	test(t, eraseOverlapIntervals3)
}

func test(t *testing.T, function func(intervals [][]int) int) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			1,
		},
		{
			"",
			args{[][]int{{1, 2}, {1, 2}, {1, 2}}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
