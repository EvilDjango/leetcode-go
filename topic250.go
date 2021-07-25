package leetcode_go

/*
250. 统计同值子树
给定一个二叉树，统计该二叉树数值相同的子树个数。

同值子树是指该子树的所有节点都拥有相同的数值。

示例：

输入: root = [5,1,5,5,5,null,5]

              5
             / \
            1   5
           / \   \
          5   5   5

输出: 4
通过次数4,037提交次数6,384

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/23/21 10:46 AM
*/

// 递归
func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	isUnival(root, &count)
	return count
}

func isUnival(root *TreeNode, count *int) bool {
	if root == nil {
		return true
	}
	left, right := root.Left, root.Right
	leftUni := isUnival(left, count)
	rightUni := isUnival(right, count)
	if (left == nil || left.Val == root.Val) && (right == nil || right.Val == root.Val) && leftUni && rightUni {
		*count++
		return true
	}
	return false
}

// 循环，后序遍历
func countUnivalSubtrees2(root *TreeNode) int {
	ans := 0
	var stack []*TreeNode
	var last *TreeNode
	universal := map[*TreeNode]bool{nil: true}
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			left, right := root.Left, root.Right
			if right == nil || right == last {
				if universal[left] && universal[right] && (left == nil || left.Val == root.Val) && (right == nil || right.Val == root.Val) {
					ans++
					universal[root] = true
				}
				stack = stack[:len(stack)-1]
				last = root
				root = nil
			} else {
				root = right
			}
		}
	}
	return ans
}

// 递归2
func countUnivalSubtrees3(root *TreeNode) int {
	count := 0
	isValidPart(root, 0, &count)
	return count
}

// 判断当前的子树能否与其父节点共同组成一个单值子树
func isValidPart(root *TreeNode, ancestor int, count *int) bool {
	if root == nil {
		return true
	}
	left, right := root.Left, root.Right
	leftUni := isValidPart(left, root.Val, count)
	rightUni := isValidPart(right, root.Val, count)
	if !leftUni || !rightUni {
		return false
	}
	*count++
	return root.Val == ancestor
}
