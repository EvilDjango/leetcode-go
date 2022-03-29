// 剑指 Offer II 007. 数组中和为 0 的三个数
//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且 不重复 的三元组。
//
//
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//示例 2：
//
//输入：nums = []
//输出：[]
//示例 3：
//
//输入：nums = [0]
//输出：[]
//
//
//提示：
//
//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105
//
//
//注意：本题与主站 15 题相同：https://leetcode-cn.com/problems/3sum/
//
//通过次数21,174提交次数47,645
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/25 上午11:31
package _fGaJU

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				if left == i+1 || nums[left] != nums[left-1] {
					ans = append(ans, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--
			}
		}
	}
	return ans
}
