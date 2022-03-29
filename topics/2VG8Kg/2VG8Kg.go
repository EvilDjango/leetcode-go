// 剑指 Offer II 008. 和大于等于 target 的最短子数组
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
//示例 1：
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//示例 2：
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//示例 3：
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//提示：
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 105
//
//
//进阶：
//
//如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
//
//
//注意：本题与主站 209 题相同：https://leetcode-cn.com/problems/minimum-size-subarray-sum/
//
//通过次数21,025提交次数42,692
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/25 上午10:46
package _VG8Kg

func minSubArrayLen(target int, nums []int) int {
	left, sum, min := 0, 0, len(nums)+1
	for right, num := range nums {
		sum += num
		for sum-nums[left] >= target {
			sum -= nums[left]
			left++
		}
		if sum >= target && right-left+1 < min {
			min = right - left + 1
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}
