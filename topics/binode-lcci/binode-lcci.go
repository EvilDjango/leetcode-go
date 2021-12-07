// 面试题 17.12. BiNode
//二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质，转换操作应是原址的，也就是在原始的二叉搜索树上直接修改。
//
//返回转换后的单向链表的头节点。
//
//注意：本题相对原题稍作改动
//
//
//
//示例：
//
//输入： [4,2,5,1,3,null,6,0]
//输出： [0,null,1,null,2,null,3,null,4,null,5,null,6]
//提示：
//
//节点数量不会超过 100000。
//通过次数20,356提交次数32,233
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/17/21 1:51 PM
package binode_lcci

import (
	"leetcode-go/container"
)

// 递归解法
func convertBiNode(root *TreeNode) *TreeNode {
	head, _ := dfs(root)
	return head
}

func dfs(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	head, tail := root, root
	left, right := root.Left, root.Right
	if left != nil {
		leftHead, LeftTail := dfs(left)
		root.Left = nil
		head = leftHead
		LeftTail.Right = root
	}
	if right != nil {
		rightHead, rightTail := dfs(right)
		tail = rightTail
		root.Right = rightHead
	}
	return head, tail
}

// 递归解法2
func convertBiNode2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	tail := dummy
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		root.Left = nil
		tail.Right = root
		tail = root
		dfs(root.Right)
	}
	dfs(root)
	return dummy.Right
}

// 循环解法
func convertBiNode3(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	dummy := &TreeNode{}
	tail := dummy
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			tail.Right = node
			tail = node
			node.Left = nil
			stack = stack[:len(stack)-1]
			root = node.Right
		}
	}
	return dummy.Right
}
