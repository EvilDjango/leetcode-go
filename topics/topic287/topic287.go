package topic287

/*
287. 寻找重复数
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。



示例 1：

输入：nums = [1,3,4,2,2]
输出：2
示例 2：

输入：nums = [3,1,3,4,2]
输出：3
示例 3：

输入：nums = [1,1]
输出：1
示例 4：

输入：nums = [1,1,2]
输出：1


提示：

1 <= n <= 105
nums.length == n + 1
1 <= nums[i] <= n
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次


进阶：

如何证明 nums 中至少存在一个重复的数字?
你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？
通过次数167,966提交次数253,148

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/13/21 2:49 PM
*/

// 参考题解，二分查找
func findDuplicate(nums []int) int {
	l, r := 1, len(nums)
	for l < r {
		mid := (r-l)/2 + l
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 二进制
// 若没有重复的数，那么 是 [0,1,3...n]
// 将所有数字二进制展开。对于某一个二进制位，记没有重复的情况下该位上1的数量为x，记实际数量为y，若y>x，那么重复数字在该二级制位上为1，否则为0
func findDuplicate2(nums []int) int {
	ans := 0
	mask := 1
	for mask <= len(nums)-1 {
		noRepeat, cnt := 0, 0
		for i, v := range nums {
			if mask&i > 0 {
				noRepeat++
			}
			if mask&v > 0 {
				cnt++
			}
		}
		if cnt > noRepeat {
			ans |= mask
		}
		mask <<= 1
	}

	return ans
}

// Floyd判圈算法
func findDuplicate3(nums []int) int {
	slow, fast := 0, 0
	for slow == 0 || slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
