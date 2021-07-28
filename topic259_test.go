package leetcode_go

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 1:15 PM
*/

func Test_threeSumSmaller2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{0, -4, -1, 1, -1, 2}, -5},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumSmaller2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumSmaller2() = %v, want %v", got, tt.want)
			}
		})
	}
}
