package topic0257

import (
	"bytes"
	"leetcode-go/tree"
	"strconv"
)

/*
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
通过次数128,042提交次数188,795

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/27/21 8:36 PM
*/

// dfs
func binaryTreePaths(root *tree.TreeNode) []string {
	var ans []string
	findPath(root, bytes.Buffer{}, &ans)
	return ans
}

func findPath(root *tree.TreeNode, buffer bytes.Buffer, paths *[]string) {
	if root == nil {
		return
	}
	buffer.WriteString(strconv.Itoa(root.Val))
	left, right := root.Left, root.Right
	if left == nil && right == nil {
		*paths = append(*paths, buffer.String())
	} else {
		size := buffer.Len()
		buffer.WriteString("->")
		findPath(left, buffer, paths)
		buffer.Truncate(size)
		findPath(right, buffer, paths)
		buffer.Truncate(size)
	}
}

// 广度优先遍历
func binaryTreePaths2(root *tree.TreeNode) []string {
	type myNode struct {
		cur       *tree.TreeNode
		ancestors string
	}
	var ans []string
	if root == nil {
		return ans
	}
	var queue = []myNode{{root, ""}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		cur := node.cur
		node.ancestors += strconv.Itoa(cur.Val)
		if cur.Left == nil && cur.Right == nil {
			ans = append(ans, node.ancestors)
			continue
		}
		node.ancestors += "->"
		if cur.Left != nil {
			queue = append(queue, myNode{cur.Left, node.ancestors})
		}

		if cur.Right != nil {
			queue = append(queue, myNode{cur.Right, node.ancestors})
		}
	}
	return ans
}
