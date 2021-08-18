package topic9527

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/3/21 8:47 PM
*/
func Test_hasLoop(t *testing.T) {
	test9527(t, hasLoop)
}

func Test_hasLoop2(t *testing.T) {
	test9527(t, hasLoop2)
}

func test9527(t *testing.T, function func(n int, edges [][]int) bool) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{5, [][]int{{0, 1}, {2, 3}}},
			false,
		},
		{
			"",
			args{5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}},
			false,
		},
		{
			"",
			args{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("validTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
