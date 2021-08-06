package leetcode_go

import (
	"leetcode-go/common"
	"sort"
)

/*
259. 较小的三数之和
给定一个长度为 n 的整数数组和一个目标值 target，寻找能够使条件 nums[i] + nums[j] + nums[k] < target 成立的三元组  i, j, k 个数（0 <= i < j < k < n）。

示例：

输入: nums = [-2,0,1,3], target = 2
输出: 2
解释: 因为一共有两个三元组满足累加和小于 2:
     [-2,0,1]
     [-2,0,3]
进阶：是否能在 O(n2) 的时间复杂度内解决？

通过次数6,318提交次数11,060
Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 10:59 AM
*/
func threeSumSmaller(nums []int, target int) int {
	n := len(nums)
	sort.Sort(common.Ints(nums))
	ans := 0
loop1:
	for i := 0; i < n; i++ {
	loop2:
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] < target {
					ans++
					continue
				}
				if k == i+2 {
					break loop1
				}
				if k == j+1 {
					break loop2
				}
				break
			}
		}
	}
	return ans
}

// 使用二分查找
func threeSumSmaller2(nums []int, target int) int {
	n := len(nums)
	sort.Sort(common.Ints(nums))
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans += common.LowerBoundPart(nums, float64(target-nums[i]-nums[j]), j+1, n) - j - 1
		}
	}
	return ans
}

// 双指针法
func threeSumSmaller3(nums []int, target int) int {
	n := len(nums)
	sort.Sort(common.Ints(nums))
	ans := 0
	for i := 0; i < n; i++ {
		j, k := i+1, n-1
		for j < k {
			if nums[j]+nums[k]+nums[i] < target {
				ans += k - j
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
