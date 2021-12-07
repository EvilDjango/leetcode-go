package topic0272

import (
	"leetcode-go/common"
	"leetcode-go/tree"
	"math"
)

/*
272. 最接近的二叉搜索树值 II
给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的 k 个值。

注意：

给定的目标值 target 是一个浮点数
你可以默认 k 值永远是有效的，即 k ≤ 总结点数
题目保证该二叉搜索树中只会存在一种 k 个值集合最接近目标值
示例：

输入: root = [4,2,5,1,3]，目标值 = 3.714286，且 k = 2

    4
   / \
  2   5
 / \
1   3

输出: [4,3]
拓展：
假设该二叉搜索树是平衡的，请问您是否能在小于 O(n)（n 为总结点数）的时间复杂度内解决该问题呢？

通过次数3,102提交次数4,832

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/5/21 10:58 AM
*/

// 中序遍历得到上升序列，双指针法从最接近target的地方依次向两边扩展
func closestKValues(root *tree.TreeNode, target float64, k int) []int {
	// 首先进行中序遍历得到升序数组
	var nums []int
	var stack []*tree.TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nums = append(nums, node.Val)
			root = node.Right
		}
	}

	// 二分查找得到下界
	high := common.LowerBound(nums, target)
	low := high - 1
	var ans []int
	// 双指针依次在两侧寻找最靠近的数字
	for len(ans) < k {
		if high < len(nums) && (low < 0 || math.Abs(float64(nums[high])-target) < math.Abs(float64(nums[low])-target)) {
			ans = append(ans, nums[high])
			high++
		} else {
			ans = append(ans, nums[low])
			low--
		}
	}
	return ans
}

// 中序遍历的同时进行收集k个元素的任务
// 显然k个元素一定是连续的，那么我们可以维护一个长度为k的滑动窗口，当窗口中的元素少于k时，扩展窗口。元素数量已经等于k时，
//如果下一个元素比窗口最左端的元素更靠近target，窗口右移一位，否则说明已经找到了最靠近的k个元素
func closestKValues2(root *tree.TreeNode, target float64, k int) []int {
	var ans []int
	var stack []*tree.TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(ans) < k {
				ans = append(ans, root.Val)
			} else if math.Abs(float64(root.Val)-target) < math.Abs(float64(ans[0])-target) {
				ans = ans[1:]
				ans = append(ans, root.Val)
			} else {
				break
			}
			root = root.Right
		}
	}
	return ans
}
