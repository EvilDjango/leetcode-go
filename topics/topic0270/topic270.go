package topic0270

import (
	"leetcode-go/topics"
	"math"
)

/*
270. 最接近的二叉搜索树值
给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。

注意：

给定的目标值 target 是一个浮点数
题目保证在该二叉搜索树中只会存在一个最接近目标值的数
示例：

输入: root = [4,2,5,1,3]，目标值 target = 3.714286

    4
   / \
  2   5
 / \
1   3

输出: 4
通过次数8,732提交次数15,929

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/4/21 4:57 PM
*/
// dfs
func closestValue(root *topics.TreeNode, target float64) int {
	v := float64(root.Val)
	diff := math.Abs(target - v)
	if v > target {
		if root.Left != nil {
			lAns := closestValue(root.Left, target)
			if math.Abs(target-float64(lAns)) < diff {
				return lAns
			}
		}
	} else if v < target {
		if root.Right != nil {
			rAns := closestValue(root.Right, target)
			if math.Abs(target-float64(rAns)) < diff {
				return rAns
			}
		}
	}
	return root.Val
}

// dfs
func closestValue2(root *topics.TreeNode, target float64) int {
	ans := root.Val
	var dfs func(root *topics.TreeNode)
	dfs = func(root *topics.TreeNode) {
		if root == nil {
			return
		}
		v := float64(root.Val)

		if math.Abs(target-v) < math.Abs(target-float64(ans)) {
			ans = root.Val
		}
		if v > target {
			dfs(root.Left)
		} else if v < target {
			dfs(root.Right)
		} else {
			return
		}
	}
	dfs(root)
	return ans
}

// 迭代，中序遍历
func closestValue3(root *topics.TreeNode, target float64) int {
	var stack []*topics.TreeNode
	cur, prev := root.Val, root.Val
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prev = cur
			cur = node.Val
			if float64(cur) >= target {
				break
			}
			root = node.Right
		}
	}
	if math.Abs(float64(prev)-target) < math.Abs(float64(cur)-target) {
		return prev
	}
	return cur
}

// 迭代，二分查找
func closestValue4(root *topics.TreeNode, target float64) int {
	ans := root.Val
	for root != nil {
		v := float64(root.Val)
		if math.Abs(v-target) < math.Abs(float64(ans)-target) {
			ans = root.Val
		}
		if v < target {
			root = root.Right
		} else if v > target {
			root = root.Left
		} else {
			break
		}
	}
	return ans
}
