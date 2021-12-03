// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/1/21 10:12 AM
package first_common_ancestor_lcci

import (
	"leetcode-go/tree"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	test(t, lowestCommonAncestor)
}

func Test_lowestCommonAncestor2(t *testing.T) {
	test(t, lowestCommonAncestor2)
}

func test(t *testing.T, function func(root *tree.TreeNode, p *tree.TreeNode, q *tree.TreeNode) *tree.TreeNode) {
	root := tree.New(3, 5, 1, 6, 2, 0, 8, tree.None, tree.None, 7, 4)
	p := root.GetChild(5)
	q := root.GetChild(4)
	type args struct {
		root *tree.TreeNode
		p    *tree.TreeNode
		q    *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *tree.TreeNode
	}{
		{
			"",
			args{root, p, q},
			p,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
