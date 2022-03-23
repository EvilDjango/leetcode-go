// 剑指 Offer II 011. 0 和 1 个数相同的子数组
//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
//示例 1:
//
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
//示例 2:
//
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
//
//
//提示：
//
//1 <= nums.length <= 105
//nums[i] 不是 0 就是 1
//
//
//注意：本题与主站 525 题相同： https://leetcode-cn.com/problems/contiguous-array/
//
//通过次数14,501提交次数25,846
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/23 下午12:12
package A1NYOS

// 暴力枚举,会超时
func findMaxLength(nums []int) int {
	n := len(nums)
	maxLen := 0
	for i := 0; i < n; i++ {
		var sum int
		if nums[i] == 1 {
			sum = 1
		} else {
			sum = -1
		}
		for j := i + 1; j < n; j++ {
			if nums[j] == 1 {
				sum += 1
			} else {
				sum += -1
			}
			if sum == 0 && j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}

// 前缀和+哈希表
func findMaxLength2(nums []int) int {
	prefix := 0
	// 初始值表示空前缀和为0
	minIndex := map[int]int{0: -1}
	maxLen := 0
	for i, num := range nums {
		if num == 1 {
			prefix += 1
		} else {
			prefix -= 1
		}
		if left, ok := minIndex[prefix]; ok {
			if i-left > maxLen {
				maxLen = i - left
			}
		} else {
			minIndex[prefix] = i
		}
	}
	return maxLen
}
