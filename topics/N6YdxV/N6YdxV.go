// 剑指 Offer II 068. 查找插入位置
//给定一个排序的整数数组 nums 和一个整数目标值 target ，请在数组中找到 target ，并返回其下标。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。
//
//
//
//示例 1:
//
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//示例 2:
//
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//示例 3:
//
//输入: nums = [1,3,5,6], target = 7
//输出: 4
//示例 4:
//
//输入: nums = [1,3,5,6], target = 0
//输出: 0
//示例 5:
//
//输入: nums = [1], target = 0
//输出: 0
//
//
//提示:
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums 为无重复元素的升序排列数组
//-104 <= target <= 104
//
//
//注意：本题与主站 35 题相同： https://leetcode-cn.com/problems/search-insert-position/
//
//通过次数5,558提交次数11,094
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/28/21 11:07 PM
package N6YdxV

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (r-l)>>1 + l
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
