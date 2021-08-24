package topic300

import (
	"leetcode-go/common"
)

/*
300. 最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1


提示：

1 <= nums.length <= 2500
-104 <= nums[i] <= 104


进阶：

你可以设计时间复杂度为 O(n2) 的解决方案吗？
你能将算法的时间复杂度降低到 O(n log(n)) 吗?
通过次数330,465提交次数647,627

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/23/21 10:05 PM
*/

// 动态规划
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	ans := 1
	for i := 0; i < n; i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				max = common.Max(max, dp[j]+1)
			}
		}
		ans = common.Max(ans, max)
		dp[i] = max
	}
	return ans
}

// 贪心 + 二分查找
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	min := make([]int, n+1)
	min[1] = nums[0]
	ans := 1
	for i := 1; i < n; i++ {
		l, r := 1, ans
		for l <= r {
			mid := (r-l)/2 + l
			if min[mid] >= nums[i] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		min[l] = nums[i]
		if l > ans {
			ans = l
		}
	}
	return ans
}
