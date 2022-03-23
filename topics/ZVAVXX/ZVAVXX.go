// 剑指 Offer II 009. 乘积小于 K 的子数组
//给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。
//
//
//
//示例 1:
//
//输入: nums = [10,5,2,6], k = 100
//输出: 8
//解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
//需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
//示例 2:
//
//输入: nums = [1,2,3], k = 0
//输出: 0
//
//
//提示:
//
//1 <= nums.length <= 3 * 104
//1 <= nums[i] <= 1000
//0 <= k <= 106
//
//
//注意：本题与主站 713 题相同：https://leetcode-cn.com/problems/subarray-product-less-than-k/
//
//通过次数17,584提交次数32,409
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/23 下午1:35
package ZVAVXX

// 滑动窗口,比较复杂的解法，首先求出各个最长的子序列，然后分别计算各个最长子序列的子序列个数，同时还要去掉重合部分的计数。
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	left := 0
	product := 1
	count := 0
	overlap := 0
	for i, num := range nums {
		product *= num
		overflow := product >= k
		if overflow {
			length := i - left
			count += length * (length + 1) / 2
			count -= overlap * (overlap + 1) / 2
		}

		for product >= k {
			product /= nums[left]
			left++
		}
		if overflow {
			overlap = i - left
		}
	}
	length := len(nums) - left
	count += length * (length + 1) / 2
	count -= overlap * (overlap + 1) / 2
	return count
}

// 滑动窗口，简便解法。每次只考虑以当前节点结尾的子序列。
func numSubarrayProductLessThanK2(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	product, left, count := 1, 0, 0
	for i, num := range nums {
		product *= num
		for product >= k {
			product /= nums[left]
			left++
		}
		count += i - left + 1
	}
	return count
}
