// 剑指 Offer II 045. 二叉树最底层最左边的值
//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
//假设二叉树中至少有一个节点。
//
//
//
//示例 1:
//
//
//
//输入: root = [2,1,3]
//输出: 1
//示例 2:
//
//
//
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
//
//
//提示:
//
//二叉树的节点个数的范围是 [1,104]
//-231 <= Node.val <= 231 - 1
//
//
//注意：本题与主站 513 题相同： https://leetcode-cn.com/problems/find-bottom-left-tree-value/
//
//通过次数5,822提交次数7,231
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/3/22 2:14 PM
package LwUNpT

import . "leetcode-go/container/tree"

func findBottomLeftValue(root *TreeNode) int {
	maxDepth, val := 0, 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth >= maxDepth {
			maxDepth, val = depth, root.Val
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 0)
	return val
}
