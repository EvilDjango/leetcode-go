// 剑指 Offer II 051. 节点之和最大的路径
//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
//路径和 是路径中各节点值的总和。
//
//给定一个二叉树的根节点 root ，返回其 最大路径和，即所有路径上节点值之和的最大值。
//
//
//
//示例 1：
//
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//示例 2：
//
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//提示：
//
//树中节点数目范围是 [1, 3 * 104]
//-1000 <= Node.val <= 1000
//
//
//注意：本题与主站 124 题相同： https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
//
//通过次数3,428提交次数8,250
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/2/22 4:51 PM
package jC7MId

import (
	. "leetcode-go/container/tree"
)

func maxPathSum(root *TreeNode) int {
	ans := root.Val
	var dfs func(node *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(0, dfs(root.Left))
		r := max(0, dfs(root.Right))
		ans = max(ans, l+r+root.Val)
		return root.Val + max(l, r)
	}
	dfs(root)
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
