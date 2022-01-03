// 剑指 Offer II 047. 二叉树剪枝
//给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
//
//节点 node 的子树为 node 本身，以及所有 node 的后代。
//
//
//
//示例 1:
//
//输入: [1,null,0,0,1]
//输出: [1,null,0,null,1]
//解释:
//只有红色节点满足条件“所有不包含 1 的子树”。
//右图为返回的答案。
//
//
//示例 2:
//
//输入: [1,0,1,0,0,0,1]
//输出: [1,null,1,null,1]
//解释:
//
//
//示例 3:
//
//输入: [1,1,0,1,1,0,1,0]
//输出: [1,1,0,1,1,null,1]
//解释:
//
//
//
//
//提示:
//
//二叉树的节点个数的范围是 [1,200]
//二叉树节点的值只会是 0 或 1
//
//
//注意：本题与主站 814 题相同：https://leetcode-cn.com/problems/binary-tree-pruning/
//
//通过次数5,440提交次数7,880
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/3/22 12:12 AM
package pOCWxh

import . "leetcode-go/container/tree"

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left, right := pruneTree(root.Left), pruneTree(root.Right)
	if root.Val == 0 && left == nil && right == nil {
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}
