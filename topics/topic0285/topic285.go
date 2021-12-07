package topic0285

import (
	"leetcode-go/tree"
)

/*
285. 二叉搜索树中的中序后继
给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。

节点 p 的后继是值比 p.val 大的节点中键值最小的节点。



示例 1：



输入：root = [2,1,3], p = 1
输出：2
解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
示例 2：



输入：root = [5,3,6,2,4,null,null,1], p = 6
输出：null
解释：因为给出的节点没有中序后继，所以答案就返回 null 了。


提示：

树中节点的数目在范围 [1, 104] 内。
-105 <= Node.val <= 105
树中各节点的值均保证唯一。
通过次数7,342提交次数11,540

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/13/21 11:07 AM
*/

func inorderSuccessor(root *tree.TreeNode, p *tree.TreeNode) *tree.TreeNode {
	var prev *tree.TreeNode
	var stack []*tree.TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if prev == p {
				break
			}
			prev = root
			root = root.Right
		}
	}
	return root
}

// 利用二叉搜索树进行二分搜索
func inorderSuccessor2(root *tree.TreeNode, p *tree.TreeNode) *tree.TreeNode {
	if p.Right != nil {
		left := p.Right
		for left.Left != nil {
			left = left.Left
		}
		return left
	}
	var prev *tree.TreeNode
	for root != p {
		if root.Val > p.Val {
			prev = root
		} else {
			root = root.Right
		}
	}
	return prev
}
