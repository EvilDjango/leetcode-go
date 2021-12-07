// 面试题 04.04. 检查平衡性
//实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。
//
//
//示例 1:
//给定二叉树 [3,9,20,null,null,15,7]
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回 true 。
//示例 2:
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//      1
//     / \
//    2   2
//   / \
//  3   3
// / \
//4   4
//返回 false 。
//通过次数34,302提交次数58,535
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/1/21 3:02 PM
package check_balance_lcci

import (
	"leetcode-go/container/tree"
	"math"
)

func isBalanced(root *tree.TreeNode) bool {
	balanced := true
	var getHeight func(root *tree.TreeNode) int
	getHeight = func(root *tree.TreeNode) int {
		if root == nil || !balanced {
			return 0
		}
		leftH := getHeight(root.Left)
		rightH := getHeight(root.Right)
		if math.Abs(float64(leftH-rightH)) > 1 {
			balanced = false
		}
		if leftH > rightH {
			return leftH
		}
		return rightH
	}
	getHeight(root)
	return balanced
}
