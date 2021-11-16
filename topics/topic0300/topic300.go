package topic0300

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

// 深度遍历+记忆化
func lengthOfLIS3(nums []int) int {
	// 从各个位置出发能得到的最大递增子序列长度
	maxLens := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dfs(nums, maxLens, i)
	}
	maxLen := 0
	for _, length := range maxLens {
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func dfs(nums []int, maxLens []int, i int) int {
	if maxLens[i] > 0 {
		return maxLens[i]
	}
	maxLen := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[j] <= nums[i] {
			continue
		}
		maxLen = max(maxLen, dfs(nums, maxLens, j))
	}
	maxLens[i] = maxLen + 1
	return maxLen + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 动态规划
func lengthOfLIS4(nums []int) int {
	n := len(nums)
	// dp[i]表示从下标i开始的最长递增子序列长度
	dp := make([]int, n)
	maxLen := 0
	for i := n - 1; i >= 0; i-- {
		forwardLen := 0
		for j := i + 1; j < n; j++ {
			if nums[j] <= nums[i] {
				continue
			}
			forwardLen = max(forwardLen, dp[j])
		}
		dp[i] = forwardLen + 1
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

// 贪心算法
func lengthOfLIS5(nums []int) int {
	n := len(nums)
	// minTail[i]表示长度为i的子序列末尾的最小值
	minTail := make([]int, n+1)
	minTail[1] = nums[0]
	maxLen := 1
	for i := 1; i < n; i++ {
		length := search(minTail, 1, maxLen, nums[i]) + 1
		minTail[length] = nums[i]
		maxLen = max(maxLen, length)
	}
	return maxLen
}

// 在nums的 [l,r)区间搜索最后一个小于target的元素返回其下标
func search(nums []int, l, r, target int) int {
	for l <= r {
		mid := (r-l)>>1 + l
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
