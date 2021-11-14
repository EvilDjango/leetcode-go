// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/13/21 5:10 PM
package search_rotate_array_lcci

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5},
			8,
		},
		{
			"",
			args{[]int{1, 1, 1, 1, 1, 2, 1, 1, 1}, 2},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
