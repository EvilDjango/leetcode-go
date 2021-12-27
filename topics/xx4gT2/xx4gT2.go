// 剑指 Offer II 076. 数组中的第 k 大的数字
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
//示例 1:
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//示例 2:
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//
//提示：
//
//1 <= k <= nums.length <= 104
//-104 <= nums[i] <= 104
//
//
//注意：本题与主站 215 题相同： https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
//
//通过次数7,774提交次数11,380
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/25/21 12:25 AM
package xx4gT2

import "math/rand"

// 利用快速排序的思想
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	l, r := 0, n-1
	for {
		if l == r {
			return nums[l]
		}
		left, right := l, r
		randI := rand.Intn(r-l) + l
		nums[l], nums[randI] = nums[randI], nums[l]
		base := nums[l]
		for l < r {
			for l < r && nums[r] >= base {
				r--
			}
			nums[l] = nums[r]
			for l < r && nums[l] <= base {
				l++
			}
			nums[r] = nums[l]
		}
		nums[l] = base
		if l < n-k {
			l, r = l+1, right
		} else if l > n-k {
			l, r = left, l-1
		} else {
			return nums[l]
		}
	}
}
