// 面试题 10.11. 峰与谷
//在一个整数数组中，“峰”是大于或等于相邻整数的元素，相应地，“谷”是小于或等于相邻整数的元素。例如，在数组{5, 8, 4, 2, 3, 4, 6}中，{8, 6}是峰， {5, 2}是谷。现在给定一个整数数组，将该数组按峰与谷的交替顺序排序。
//
//示例:
//
//输入: [5, 3, 1, 2, 3]
//输出: [5, 1, 3, 2, 3]
//提示：
//
//nums.length <= 10000
//通过次数7,994提交次数11,991
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/9/21 10:09 PM
package peaks_and_valleys_lcci

import "sort"

// 先将数组进行递增排序，然后将偶数位与下一个元素交换位置，时间复杂度为nlgn
func wiggleSort(nums []int) {
	sort.Ints(nums)
	for i := 0; i+1 < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}

// 一次遍历，依次判断每个位置是否符合要求，不符合要求则交换
func wiggleSort2(nums []int) {
	if len(nums) < 3 {
		return
	}
	// 当前是否为峰
	up := nums[1] >= nums[0]
	for i := 1; i+1 < len(nums); i++ {
		if (up && nums[i] < nums[i+1]) || (!up && nums[i] > nums[i+1]) {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
		up = !up
	}
}
