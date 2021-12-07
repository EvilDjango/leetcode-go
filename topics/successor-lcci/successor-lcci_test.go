// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/1/21 1:19 PM
package successor_lcci

import (
	"leetcode-go/topics"
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	test(t, inorderSuccessor)
}
func Test_inorderSuccessor2(t *testing.T) {
	test(t, inorderSuccessor2)
}
func Test_inorderSuccessor3(t *testing.T) {
	test(t, inorderSuccessor3)
}
func Test_inorderSuccessor4(t *testing.T) {
	test(t, inorderSuccessor4)
}

func test(t *testing.T, function func(root *topics.TreeNode, p *topics.TreeNode) *topics.TreeNode) {
	root := topics.New(2, topics.None, 3)
	p := root.GetChild(2)
	target := root.GetChild(3)
	type args struct {
		root *topics.TreeNode
		p    *topics.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *topics.TreeNode
	}{
		{
			"",
			args{root, p},
			target,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
