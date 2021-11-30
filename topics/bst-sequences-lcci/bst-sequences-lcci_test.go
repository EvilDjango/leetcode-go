// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/30/21 3:35 PM
package bst_sequences_lcci

import (
	"leetcode-go/common"
	"leetcode-go/tree"
	"testing"
)

func TestBSTSequences(t *testing.T) {
	test(t, BSTSequences)
}

func TestBSTSequences2(t *testing.T) {
	test(t, BSTSequences2)
}

func TestBSTSequences3(t *testing.T) {
	test(t, BSTSequences3)
}

func test(t *testing.T, function func(root *tree.TreeNode) [][]int) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"",
			args{tree.New(1, 2)},
			[][]int{{1, 2}},
		},
		{
			"",
			args{tree.New(2, 1, 3)},
			[][]int{{2, 3, 1}, {2, 1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.root); !common.TwoDimensionArrayEqualsIgnoreOrder(got, tt.want) {
				t.Errorf("BSTSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
