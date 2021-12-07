// 面试题 04.05. 合法二叉搜索树
//实现一个函数，检查一棵二叉树是否为二叉搜索树。
//
//示例 1:
//输入:
//    2
//   / \
//  1   3
//输出: true
//示例 2:
//输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//通过次数26,758提交次数75,887
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/1/21 2:51 PM
package legal_binary_search_tree_lcci

import (
	"leetcode-go/topics"
	"math"
)

func isValidBST(root *topics.TreeNode) bool {
	return isValid(root, math.MaxInt64, math.MinInt64)
}

func isValid(root *topics.TreeNode, ceiling, floor int) bool {
	if root == nil {
		return true
	}
	if root.Val >= ceiling || root.Val <= floor {
		return false
	}
	return isValid(root.Left, root.Val, floor) && isValid(root.Right, ceiling, root.Val)
}
