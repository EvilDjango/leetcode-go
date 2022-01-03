// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/31/21 11:58 AM
package g5c51o

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	test(t, topKFrequent)
}

func Test_topKFrequent2(t *testing.T) {
	test(t, topKFrequent2)
}

func test(t *testing.T, f func(nums []int, k int) []int) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{[]int{4, 1, -1, 2, -1, 2, 3}, 2},
			[]int{-1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
