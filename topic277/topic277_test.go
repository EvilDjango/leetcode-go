package topic277

import (
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/10/21 7:43 PM
*/

func Test_solution(t *testing.T) {
	type args struct {
		matrix [][]int
		n      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				[][]int{
					{0, 1},
					{1, 0},
				},
				2,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if celebrity := solution(func(a int, b int) bool {
				return tt.args.matrix[a][b] == 1
			})(tt.args.n); celebrity != tt.want {
				t.Errorf("solution() = %v, want %v", celebrity, tt.want)
			}
		})
	}
}
