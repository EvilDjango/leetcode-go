package topic261

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/3/21 8:24 PM
*/

func Test_validTree(t *testing.T) {
	test261(t, validTree)
}

func Test_validTree2(t *testing.T) {
	test261(t, validTree2)
}

func Test_validTree3(t *testing.T) {
	test261(t, validTree3)
}

func test261(t *testing.T, function func(n int, edges [][]int) bool) {
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
			true,
		},
		{
			"",
			args{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}},
			false,
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
