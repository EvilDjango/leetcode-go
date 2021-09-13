package topic0298

import (
	"leetcode-go/common"
	"leetcode-go/tree"
)

/*
298. 二叉树最长连续序列
给你一棵指定的二叉树的根节点 root ，请你计算其中 最长连续序列路径 的长度。

最长连续序列路径 是依次递增 1 的路径。该路径，可以是从某个初始节点到树中任意节点，通过「父 - 子」关系连接而产生的任意路径。且必须从父节点到子节点，反过来是不可以的。


示例 1：


输入：root = [1,null,3,2,4,null,null,null,5]
输出：3
解释：当中，最长连续序列是 3-4-5 ，所以返回结果为 3 。
示例 2：


输入：root = [2,null,3,2,null,1]
输出：2
解释：当中，最长连续序列是 2-3 。注意，不是 3-2-1，所以返回 2 。


提示：

树中节点的数目在范围 [1, 3 * 104] 内
-3 * 104 <= Node.val <= 3 * 104
通过次数4,523提交次数7,808

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/23/21 9:15 PM
*/

// 自顶向下深度遍历
func longestConsecutive(root *TreeNode) int {
	var dfs func(ans, root *TreeNode, cnt int) int
	dfs = func(ans, root *TreeNode, cnt int) int {
		if root == nil {
			return cnt
		}
		start := 1
		if ans != nil && root.Val == ans.Val+1 {
			cnt++
			start = cnt
		}
		return common.Max(cnt, common.Max(dfs(root, root.Left, start), dfs(root, root.Right, start)))
	}
	return dfs(nil, root, 0)
}

// 自底向上深度遍历
func longestConsecutive2(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	ans := 1
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		L, R := dfs(root.Left)+1, dfs(root.Right)+1
		if root.Left != nil && root.Left.Val != root.Val+1 {
			L = 1
		}
		if root.Right != nil && root.Right.Val != root.Val+1 {
			R = 1
		}
		length := Max(L, R)
		ans = Max(ans, length)
		return length
	}
	dfs(root)
	return ans
}
