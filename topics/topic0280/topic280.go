package topic0280

import "sort"

/*
280. 摆动排序
给你一个无序的数组 nums, 将该数字 原地 重排后使得 nums[0] <= nums[1] >= nums[2] <= nums[3]...。

示例:

输入: nums = [3,5,2,1,6,4]
输出: 一个可能的解答是 [3,5,1,6,2,4]
通过次数5,237提交次数7,619

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/11/21 1:59 PM
*/

func wiggleSort(nums []int) {
	sort.Sort(sort.IntSlice(nums))
	for i := 1; i+1 < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}

func wiggleSort2(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] != (i%2 == 1) {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}
