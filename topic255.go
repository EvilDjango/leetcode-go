package leetcode_go

import "math"

/*
255. 验证前序遍历序列二叉搜索树
给定一个整数数组，你需要验证它是否是一个二叉搜索树正确的先序遍历序列。

你可以假定该序列中的数都是不相同的。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [5,2,6,1,3]
输出: false
示例 2：

输入: [5,2,1,3,6]
输出: true
进阶挑战：

您能否使用恒定的空间复杂度来完成此题？

通过次数4,719提交次数9,981

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/27/21 3:44 PM
*/

func verifyPreorder(preorder []int) bool {
	var stack []int
	min := math.MinInt32
	for _, i := range preorder {
		if i < min {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] < i {
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return true
}

// 使用原数组作为栈从而不使用额外空间
func verifyPreorder2(preorder []int) bool {
	i := -1
	min := math.MinInt32
	for _, j := range preorder {
		if j < min {
			return false
		}
		for i >= 0 && preorder[i] < j {
			min = preorder[i]
			i--
		}
		i++
		preorder[i] = j
	}
	return true
}
