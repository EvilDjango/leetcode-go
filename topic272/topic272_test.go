package topic272

import (
	"leetcode-go/common"
	"leetcode-go/tree"
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/5/21 5:22 PM
*/

func Test_closestKValues(t *testing.T) {
	test_closestKValues(t, closestKValues)
}

func Test_closestKValues2(t *testing.T) {
	test_closestKValues(t, closestKValues2)
}

func test_closestKValues(t *testing.T, function func(root *tree.TreeNode, target float64, k int) []int) {
	type args struct {
		root   *tree.TreeNode
		target float64
		k      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{tree.New(4, 2, 5, 1, 3), 3.714286, 2},
			[]int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.root, tt.args.target, tt.args.k); !common.EqualIgnoreOrder(got, tt.want) {
				t.Errorf("closestKValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
