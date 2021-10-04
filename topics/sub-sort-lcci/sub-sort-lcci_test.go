// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/25/21 7:35 PM
package sub_sort_lcci

import (
	"reflect"
	"testing"
)

func Test_subSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{[]int{1, 3, 5, 7, 9}},
			[]int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
