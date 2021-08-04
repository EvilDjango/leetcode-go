package leetcode_go

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/2/21 6:57 PM
*/

func Test_canFinish(t *testing.T) {
	test(t, canFinish)
}

func Test_canFinish2(t *testing.T) {
	test(t, canFinish2)
}

func Test_canFinish3(t *testing.T) {
	test(t, canFinish3)
}

func test(t *testing.T, function func(numCourses int, prerequisites [][]int) bool) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				2,
				[][]int{
					{1, 0},
					{0, 1},
				},
			},
			false,
		},
		{
			"",
			args{
				3,
				[][]int{
					{1, 0},
					{1, 2},
					{0, 1},
				},
			},
			false,
		},
		{
			"",
			args{
				5,
				[][]int{
					{1, 4},
					{2, 4},
					{3, 1},
					{3, 2},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
